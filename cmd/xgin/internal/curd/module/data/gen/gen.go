package gen

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/module/data/template"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/output"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/parser"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/util/filex"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/util/stringx"
)

type (
	DefaultGenerator struct {
		dir       string
		pkg       string
		parentPkg string
	}

	code struct {
		importsCode string
		typesCode   string
		newCode     string
		createCode  string
		findCode    []string
		updateCode  string
		deleteCode  string
		cacheExtra  string
	}
)

// 需要增加一个对 data.go 的检测，没有得话需要创建

// NewDefaultGenerator creates an instance for defaultGenerator
func NewDefaultGenerator(dir string, parentPkg string) (*DefaultGenerator, error) {
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

	generator := &DefaultGenerator{
		dir:       dir,
		pkg:       pkg,
		parentPkg: parentPkg,
	}

	return generator, nil
}

func (g *DefaultGenerator) CreateFromTables(tables []*parser.Table) error {
	m := make(map[string]string)
	for _, e := range tables {
		code, err := g.genBiz(e)
		if err != nil {
			return err
		}
		m[e.Name.Source()] = code
	}
	return g.createFile(m)
}

func (g *DefaultGenerator) createFile(modelList map[string]string) error {
	dirAbs, err := filepath.Abs(g.dir)
	if err != nil {
		return err
	}
	g.dir = dirAbs
	g.pkg = stringx.SafeString(filepath.Base(dirAbs))
	err = filex.MkdirIfNotExist(dirAbs)
	if err != nil {
		return err
	}

	for tableName, code := range modelList {
		tn := stringx.From(tableName)
		modelFilename := fmt.Sprintf("%s", tn.Source())

		name := stringx.SafeString(modelFilename) + ".go"
		filename := filepath.Join(dirAbs, name)
		if filex.FileExists(filename) {
			log.Printf("%s already exists, ignored.", name)
			continue
		}
		err = os.WriteFile(filename, []byte(code), os.ModePerm)
		if err != nil {
			return err
		}
	}

	log.Println("data done.")
	return nil
}

func (g *DefaultGenerator) genBiz(in *parser.Table) (string, error) {
	if len(in.PrimaryKey.Name.Source()) == 0 {
		return "", fmt.Errorf("table %s: missing primary key", in.Name.Source())
	}

	importsCode, err := genImports(in, g.parentPkg)
	if err != nil {
		return "", err
	}

	createCode, err := genCreate(in)
	if err != nil {
		return "", err
	}

	findCode := make([]string, 0)
	findOneCode, err := genFindOne(in)
	if err != nil {
		return "", err
	}
	findCode = append(findCode, findOneCode)

	updateCode, err := genUpdate(in)
	if err != nil {
		return "", err
	}

	deleteCode, err := genDelete(in)
	if err != nil {
		return "", err
	}

	typesCode, err := genTypes(in)
	if err != nil {
		return "", err
	}

	newCode, err := genNew(in)
	if err != nil {
		return "", err
	}

	code := &code{
		importsCode: importsCode,
		typesCode:   typesCode,
		newCode:     newCode,
		createCode:  createCode,
		findCode:    findCode,
		updateCode:  updateCode,
		deleteCode:  deleteCode,
	}

	buffer, err := g.executeBiz(in, code)
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}

func (g *DefaultGenerator) executeBiz(table *parser.Table, code *code) (*bytes.Buffer, error) {
	text, err := filex.LoadTemplate(dataGenTemplateFile, template.ServiceGen)
	if err != nil {
		return nil, err
	}
	t := output.With("data").
		Parse(text).
		GoFmt(true)
	buffer, err := t.Execute(map[string]interface{}{
		"pkg":         g.pkg,
		"imports":     code.importsCode,
		"types":       code.typesCode,
		"new":         code.newCode,
		"create":      code.createCode,
		"find":        strings.Join(code.findCode, "\n"),
		"update":      code.updateCode,
		"delete":      code.deleteCode,
		"extraMethod": code.cacheExtra,
		"data":        table,
	})
	if err != nil {
		return nil, err
	}
	return buffer, nil
}
