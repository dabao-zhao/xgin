package gen

import "github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/util/filex"

const (
	importTemplateFile  = "type/import.tpl"
	tagTemplateFile     = "type/tag.tpl"
	fieldTemplateFile   = "type/field.tpl"
	typeTemplateFile    = "type/type.tpl"
	typeGenTemplateFile = "type/type_gen.tpl"
)

// Clean deletes all template files
func Clean() error {
	return filex.Clean()
}
