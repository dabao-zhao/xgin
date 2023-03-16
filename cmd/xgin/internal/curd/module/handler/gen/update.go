package gen

import (
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/module/handler/template"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/output"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/parser"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/util/filex"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/util/stringx"
)

func genUpdate(table *parser.Table) (string, error) {
	for _, field := range table.Fields {
		camel := stringx.SafeString(field.Name.ToCamel())
		if camel == "CreateTime" || camel == "UpdateTime" || camel == "CreateAt" || camel == "UpdateAt" {
			continue
		}

		if field.Name.Source() == table.PrimaryKey.Name.Source() {
			continue
		}

	}

	camelTableName := table.Name.ToCamel()
	text, err := filex.LoadTemplate(updateMethodTemplateFile, template.UpdateMethod)
	if err != nil {
		return "", err
	}

	methodBuffer, err := output.With("updateMethod").Parse(text).Execute(
		map[string]interface{}{
			"upperStartCamelObject": camelTableName,
			"lowerStartCamelObject": stringx.From(camelTableName).Untitle(),
			"data":                  table,
		},
	)
	if err != nil {
		return "", nil
	}

	return methodBuffer.String(), nil
}
