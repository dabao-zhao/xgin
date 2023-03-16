package gen

import (
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/module/type/template"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/output"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/parser"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/util/filex"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/util/stringx"
)

func genTypes(table *parser.Table) (string, error) {
	fields := table.Fields
	fieldsString, err := genFields(table, fields)
	if err != nil {
		return "", err
	}

	text, err := filex.LoadTemplate(typeTemplateFile, template.TypeTpl)
	if err != nil {
		return "", err
	}

	// fields 要分成单纯主键，无主键，有主键，tag 含 form，tag 不含 form 的情况
	buffer, err := output.With("types").
		Parse(text).
		Execute(map[string]interface{}{
			"upperStartCamelPrimaryKey": stringx.EscapeGolangKeyword(stringx.From(table.PrimaryKey.Name.ToCamel()).Title()),
			"lowerStartCamelPrimaryKey": stringx.EscapeGolangKeyword(stringx.From(table.PrimaryKey.Name.ToCamel()).Untitle()),
			"dataType":                  table.PrimaryKey.DataType,
			"upperStartCamelObject":     table.Name.ToCamel(),
			"fields":                    fieldsString,
			"data":                      table,
		})
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
