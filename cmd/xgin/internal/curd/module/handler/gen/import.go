package gen

import (
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/module/handler/template"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/output"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/parser"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/util/filex"
	"github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/util/packagex"
)

func genImports(table *parser.Table, parentPkg string) (string, error) {
	text, err := filex.LoadTemplate(importTemplateFile, template.ImportTpl)
	if err != nil {
		return "", err
	}

	rootPackage := packagex.GetRootPackage()
	importPackage := rootPackage + "/" + parentPkg

	buffer, err := output.With("import").Parse(text).Execute(map[string]interface{}{
		"data":            table,
		"bizPackage":      importPackage + "/biz",
		"servicesPackage": importPackage + "/services",
		"typesPackage":    importPackage + "/types",
	})
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
