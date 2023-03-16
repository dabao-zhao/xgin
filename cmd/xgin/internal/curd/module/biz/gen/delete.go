package gen

import (
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/module/biz/template"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/output"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/parser"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/util/filex"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/util/stringx"
)

func genDelete(table *parser.Table) (string, string, error) {
	camel := table.Name.ToCamel()
	text, err := filex.LoadTemplate(deleteMethodTemplateFile, template.DeleteMethod)
	if err != nil {
		return "", "", err
	}

	methodBuffer, err := output.With("delete").
		Parse(text).
		Execute(map[string]interface{}{
			"upperStartCamelObject":     camel,
			"lowerStartCamelPrimaryKey": stringx.EscapeGolangKeyword(stringx.From(table.PrimaryKey.Name.ToCamel()).Untitle()),
			"dataType":                  table.PrimaryKey.DataType,
			"data":                      table,
		})
	if err != nil {
		return "", "", err
	}

	// interface method
	text, err = filex.LoadTemplate(deleteInterfaceTemplateFile, template.DeleteMethodInterface)
	if err != nil {
		return "", "", err
	}

	interfaceBuffer, err := output.With("deleteMethod").
		Parse(text).
		Execute(map[string]interface{}{
			"lowerStartCamelPrimaryKey": stringx.EscapeGolangKeyword(stringx.From(table.PrimaryKey.Name.ToCamel()).Untitle()),
			"dataType":                  table.PrimaryKey.DataType,
			"data":                      table,
		})
	if err != nil {
		return "", "", err
	}

	return methodBuffer.String(), interfaceBuffer.String(), nil
}
