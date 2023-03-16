package gen

import "github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/util/filex"

const (
	importTemplateFile           = "biz/import.tpl"
	createTemplateMethodFile     = "biz/create_method.tpl"
	createTemplateInterfaceFile  = "biz/create_interface.tpl"
	findOneMethodTemplateFile    = "biz/find_one_method.tpl"
	findOneInterfaceTemplateFile = "biz/find_one_interface.tpl"
	updateMethodTemplateFile     = "biz/update_method.tpl"
	updateInterfaceTemplateFile  = "biz/update_interface.tpl"
	deleteMethodTemplateFile     = "biz/delete_method.tpl"
	deleteInterfaceTemplateFile  = "biz/delete_interface.tpl"
	tagTemplateFile              = "biz/tag.tpl"
	fieldTemplateFile            = "biz/field.tpl"
	typeTemplateFile             = "biz/type.tpl"
	bizNewTemplateFile           = "biz/biz_new.tpl"
	bizGenTemplateFile           = "biz/biz_gen.tpl"
)

// Clean deletes all template files
func Clean() error {
	return filex.Clean()
}
