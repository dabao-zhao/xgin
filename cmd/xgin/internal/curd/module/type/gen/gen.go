package gen

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/module/type/template"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/output"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/parser"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/util/filex"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/util/stringx"
)

type (
	DefaultGenerator struct {
		dir string
		pkg string
	}

	code struct {
		importsCode string
		typesCode   string
	}
)

// 需要优化掉 resp 的 form

// NewDefaultGenerator creates an instance for defaultGenerator
func NewDefaultGenerator(dir string) (*DefaultGenerator, error) {
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
		dir: dir,
		pkg: pkg,
	}

	return generator, nil
}

func (g *DefaultGenerator) CreateFromTables(tables []*parser.Table) error {
	m := make(map[string]string)
	for _, e := range tables {
		code, err := g.genType(e)
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

	log.Println("type done.")
	return nil
}

func (g *DefaultGenerator) genType(in *parser.Table) (string, error) {
	if len(in.PrimaryKey.Name.Source()) == 0 {
		return "", fmt.Errorf("table %s: missing primary key", in.Name.Source())
	}

	importsCode, err := genImports(in, in.ContainsTime())
	if err != nil {
		return "", err
	}

	typesCode, err := genTypes(in)
	if err != nil {
		return "", err
	}

	code := &code{
		typesCode:   typesCode,
		importsCode: importsCode,
	}

	buffer, err := g.executeType(in, code)
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}

func (g *DefaultGenerator) executeType(table *parser.Table, code *code) (*bytes.Buffer, error) {
	text, err := filex.LoadTemplate(typeGenTemplateFile, template.TypeGen)
	if err != nil {
		return nil, err
	}
	t := output.With("type").
		Parse(text).
		GoFmt(true)
	buffer, err := t.Execute(map[string]interface{}{
		"pkg":     g.pkg,
		"types":   code.typesCode,
		"imports": code.importsCode,
		"data":    table,
	})
	if err != nil {
		return nil, err
	}
	return buffer, nil
}
