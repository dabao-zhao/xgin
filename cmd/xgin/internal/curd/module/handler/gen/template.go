package gen

import "github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/util/filex"

const (
	importTemplateFile        = "handler/import.tpl"
	createTemplateMethodFile  = "handler/create_method.tpl"
	findOneMethodTemplateFile = "handler/find_one_method.tpl"
	updateMethodTemplateFile  = "handler/update_method.tpl"
	deleteMethodTemplateFile  = "handler/delete_method.tpl"
	typeTemplateFile          = "handler/type.tpl"
	handlerNewTemplateFile    = "handler/handler_new.tpl"
	handlerGenTemplateFile    = "handler/handler_gen.tpl"
)

// Clean deletes all template files
func Clean() error {
	return filex.Clean()
}
