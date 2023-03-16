package gen

import (
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/module/type/template"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/output"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/parser"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/util/filex"
)

func genTag(table *parser.Table, in string) (string, error) {
	if in == "" {
		return in, nil
	}

	text, err := filex.LoadTemplate(tagTemplateFile, template.TagTpl)
	if err != nil {
		return "", err
	}

	buffer, err := output.With("tag").Parse(text).Execute(map[string]interface{}{
		"field": in,
		"data":  table,
	})
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}

func genTagOnlyJson(table *parser.Table, in string) (string, error) {
	if in == "" {
		return in, nil
	}

	text, err := filex.LoadTemplate(tagTemplateFile, template.TagOnlyJsonTpl)
	if err != nil {
		return "", err
	}

	buffer, err := output.With("tag").Parse(text).Execute(map[string]interface{}{
		"field": in,
		"data":  table,
	})
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
