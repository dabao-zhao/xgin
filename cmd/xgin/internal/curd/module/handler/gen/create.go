package gen

import (
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/module/handler/template"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/output"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/parser"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/util/filex"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/util/stringx"
)

func genCreate(table *parser.Table) (string, error) {
	var count int
	for _, field := range table.Fields {
		camel := stringx.SafeString(field.Name.ToCamel())
		if camel == "CreateTime" || camel == "UpdateTime" || camel == "CreateAt" || camel == "UpdateAt" {
			continue
		}

		if field.Name.Source() == table.PrimaryKey.Name.Source() {
			if table.PrimaryKey.AutoIncrement {
				continue
			}
		}

		count += 1
	}

	camel := table.Name.ToCamel()
	text, err := filex.LoadTemplate(createTemplateMethodFile, template.CreateMethod)
	if err != nil {
		return "", err
	}

	methodBuffer, err := output.With("createMethod").
		Parse(text).
		Execute(map[string]interface{}{
			"upperStartCamelObject": camel,
			"lowerStartCamelObject": stringx.From(camel).Untitle(),
			"data":                  table,
		})
	if err != nil {
		return "", err
	}

	return methodBuffer.String(), nil
}
