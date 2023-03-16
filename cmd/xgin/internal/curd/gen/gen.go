package gen

import (
	"log"
	"path/filepath"

	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/model"
	bizGen "github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/module/biz/gen"
	dataGen "github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/module/data/gen"
	serviceGen "github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/module/service/gen"
	typeGen "github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/module/type/gen"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/parser"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/util/filex"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/util/stringx"
)

const pwd = "."

type (
	Generator struct {
		originDir string
		dir       string
		pkg       string
	}
)

// NewGenerator creates an instance for Generator
func NewGenerator(dir string) (*Generator, error) {
	if dir == "" {
		dir = pwd
	}
	originDir := dir
	dirAbs, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}

	dir = dirAbs
	pkg := stringx.SafeString(filepath.Base(dirAbs))
	err = filex.MkdirIfNotExist(dir)
	if err != nil {
		return nil, err
	}

	generator := &Generator{
		originDir: originDir,
		dir:       dir,
		pkg:       pkg,
	}

	return generator, nil
}

func (g *Generator) StartFromDDL(filename string, strict bool, database string) error {
	tables, err := parser.Parse(filename, database, strict)
	if err != nil {
		return err
	}
	return g.CreateFromTables(tables)
}

func (g *Generator) StartFromInformationSchema(tables map[string]*model.Table, strict bool) error {
	var m []*parser.Table
	for _, each := range tables {
		table, err := parser.ConvertDataType(each, strict)
		if err != nil {
			return err
		}
		m = append(m, table)
	}
	return g.CreateFromTables(m)
}

func (g *Generator) CreateFromTables(tables []*parser.Table) error {
	biz, err := bizGen.NewDefaultGenerator(g.dir + "/biz")
	if err != nil {
		return err
	}
	err = biz.CreateFromTables(tables)
	if err != nil {
		return err
	}

	typ, err := typeGen.NewDefaultGenerator(g.dir + "/types")
	if err != nil {
		return err
	}
	err = typ.CreateFromTables(tables)
	if err != nil {
		return err
	}

	service, err := serviceGen.NewDefaultGenerator(g.dir+"/service", g.originDir)
	if err != nil {
		return err
	}
	err = service.CreateFromTables(tables)
	if err != nil {
		return err
	}

	data, err := dataGen.NewDefaultGenerator(g.dir+"/data", g.originDir)
	if err != nil {
		return err
	}
	err = data.CreateFromTables(tables)
	if err != nil {
		return err
	}

	log.Println("done.")
	return nil
}
