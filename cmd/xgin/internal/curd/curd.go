package curd

import (
	"log"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
	gormMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/gen"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/model"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/util/filex"
)

// 初始化 curd

// CmdCurd represents the new command.
var (
	CmdCurd = &cobra.Command{
		Use:   "curd",
		Short: "Create a curd template",
		Long:  "Create a curd using the template. Example: xgin curd",
	}

	CmdCurdDDL = &cobra.Command{
		Use:   "ddl",
		Short: "From ddl create a curd template",
		Long:  "Create a curd using the template from ddl. Example: xgin curd ddl -s user.sql -d user",
		Run:   ddlRun,
	}

	CmdCurdDataSource = &cobra.Command{
		Use:   "datasource",
		Short: "From datasource create a curd template",
		Long:  "Create a curd using the template from datasource. Example: xgin curd datasource -u dsn -t * -d user",
		Run:   dataSourceRun,
	}

	src          string
	dir          string
	url          string
	tablePattern []string
)

func init() {
	CmdCurdDDL.Flags().StringVarP(&src, "src", "s", "", "The path or path globbing patterns of the ddl")
	CmdCurdDDL.Flags().StringVarP(&dir, "dir", "d", "", "The target dir")

	CmdCurdDataSource.Flags().StringVarP(&url, "url", "u", "", `The data source of database,like "root:password@tcp(127.0.0.1:3306)/database"`)
	CmdCurdDataSource.Flags().StringSliceVarP(&tablePattern, "table", "t", nil, "The table or table globbing patterns in the database")
	CmdCurdDataSource.Flags().StringVarP(&dir, "dir", "d", "", "The target dir")

	CmdCurd.AddCommand(CmdCurdDDL)
	CmdCurd.AddCommand(CmdCurdDataSource)
}

// 需要检测目录里是否有 wire

func ddlRun(cmd *cobra.Command, args []string) {
	src := strings.TrimSpace(src)
	if len(src) == 0 {
		log.Fatalln("expected path or path globbing patterns, but nothing found")
	}
	files, err := filex.Match(src)
	if err != nil {
		log.Fatalln(err)
	}
	if len(files) == 0 {
		log.Fatalln("not found any sql file")
	}

	generator, err := gen.NewGenerator(dir)
	if err != nil {
		log.Fatalln(err)
	}

	for _, f := range files {
		err = generator.StartFromDDL(f, false, "")
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func dataSourceRun(cmd *cobra.Command, args []string) {
	if len(url) == 0 {
		log.Fatalf("%v", "expected data source of mysql, but nothing found")

	}
	if len(tablePattern) == 0 {
		log.Fatalf("%v", "expected table or table globbing patterns, but nothing found")
	}

	dsn, err := mysql.ParseDSN(url)
	if err != nil {
		log.Fatalln(err)
	}

	databaseSource := strings.TrimSuffix(url, "/"+dsn.DBName) + "/information_schema"
	db, err := gorm.Open(gormMysql.Open(databaseSource), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	im := model.NewInformationSchemaModel(db)

	tables, err := im.GetAllTables(dsn.DBName)
	if err != nil {
		log.Fatalln(err)
	}

	matchTables := make(map[string]*model.Table)
	patterns := parseTableList(tablePattern)
	for _, item := range tables {
		if !patterns.Match(item) {
			continue
		}

		columnData, err := im.FindColumns(dsn.DBName, item)
		if err != nil {
			log.Fatalln(err)
		}

		table, err := columnData.Convert()
		if err != nil {
			log.Fatalln(err)
		}

		matchTables[item] = table
	}
	if len(matchTables) == 0 {
		log.Fatalln("no tables matched")
	}

	generator, err := gen.NewGenerator(dir)
	if err != nil {
		log.Fatalln(err)
	}

	err = generator.StartFromInformationSchema(matchTables, false)
	if err != nil {
		log.Fatalln(err)
	}
}
