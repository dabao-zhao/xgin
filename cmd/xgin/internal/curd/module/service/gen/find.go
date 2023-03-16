package gen

import (
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/module/service/template"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/output"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/parser"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/util/filex"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/util/stringx"
)

func genFindOne(table *parser.Table) (string, error) {
	camel := table.Name.ToCamel()
	text, err := filex.LoadTemplate(findOneMethodTemplateFile, template.FindOneMethod)
	if err != nil {
		return "", err
	}

	methodBuffer, err := output.With("findOneMethod").
		Parse(text).
		Execute(map[string]interface{}{
			"upperStartCamelObject":     camel,
			"upperStartCamelPrimaryKey": stringx.EscapeGolangKeyword(stringx.From(table.PrimaryKey.Name.ToCamel()).Title()),
			"data":                      table,
		})
	if err != nil {
		return "", err
	}

	return methodBuffer.String(), nil
}
