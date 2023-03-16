package gen

import "github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/util/filex"

const (
	importTemplateFile        = "service/import.tpl"
	createTemplateMethodFile  = "service/create_method.tpl"
	findOneMethodTemplateFile = "service/find_one_method.tpl"
	updateMethodTemplateFile  = "service/update_method.tpl"
	deleteMethodTemplateFile  = "service/delete_method.tpl"
	typeTemplateFile          = "service/type.tpl"
	serviceNewTemplateFile    = "service/service_new.tpl"
	serviceGenTemplateFile    = "service/service_gen.tpl"
)

// Clean deletes all template files
func Clean() error {
	return filex.Clean()
}
