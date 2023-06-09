package gen

import (
	"strings"

	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/module/type/template"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/output"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/parser"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/util/filex"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/util/stringx"
)

func genFields(table *parser.Table, fields []*parser.Field) (string, error) {
	var list []string

	for _, field := range fields {
		result, err := genField(table, field)
		if err != nil {
			return "", err
		}

		list = append(list, result)
	}

	return strings.Join(list, "\n"), nil
}

func genField(table *parser.Table, field *parser.Field) (string, error) {
	tag, err := genTag(table, field.NameOriginal)
	if err != nil {
		return "", err
	}

	tagOnlyJson, err := genTagOnlyJson(table, field.NameOriginal)
	if err != nil {
		return "", err
	}

	text, err := filex.LoadTemplate(fieldTemplateFile, template.FieldTpl)
	if err != nil {
		return "", err
	}

	buffer, err := output.With("types").
		Parse(text).
		Execute(map[string]interface{}{
			"name":        stringx.SafeString(field.Name.ToCamel()),
			"type":        field.DataType,
			"tag":         tag,
			"tagOnlyJson": tagOnlyJson,
			"hasComment":  field.Comment != "",
			"comment":     field.Comment,
			"data":        table,
		})
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}

func genFieldsTagOnlyJson(table *parser.Table, fields []*parser.Field) (string, error) {
	var list []string

	for _, field := range fields {
		result, err := genFieldTagOnlyJson(table, field)
		if err != nil {
			return "", err
		}

		list = append(list, result)
	}

	return strings.Join(list, "\n"), nil
}

func genFieldTagOnlyJson(table *parser.Table, field *parser.Field) (string, error) {
	tag, err := genTagOnlyJson(table, field.NameOriginal)
	if err != nil {
		return "", err
	}

	text, err := filex.LoadTemplate(fieldTemplateFile, template.FieldTpl)
	if err != nil {
		return "", err
	}

	buffer, err := output.With("types").
		Parse(text).
		Execute(map[string]interface{}{
			"name":       stringx.SafeString(field.Name.ToCamel()),
			"type":       field.DataType,
			"tag":        tag,
			"hasComment": field.Comment != "",
			"comment":    field.Comment,
			"data":       table,
		})
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
