package gen

import (
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/module/data/template"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/output"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/parser"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/util/filex"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/util/stringx"
)

func genNew(table *parser.Table) (string, error) {
	text, err := filex.LoadTemplate(dataNewTemplateFile, template.NewTpl)
	if err != nil {
		return "", err
	}

	camel := table.Name.ToCamel()
	buffer, err := output.With("new").
		Parse(text).
		Execute(map[string]interface{}{
			"upperStartCamelObject": camel,
			"lowerStartCamelObject": stringx.From(table.Name.ToCamel()).Untitle(),
			"data":                  table,
		})
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
