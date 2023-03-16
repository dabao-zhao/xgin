package gen

import "github.com/dabao-zhao/xgin/cmd/xgin/internal/curd/util/filex"

const (
	importTemplateFile        = "data/import.tpl"
	createTemplateMethodFile  = "data/create_method.tpl"
	findOneMethodTemplateFile = "data/find_one_method.tpl"
	updateMethodTemplateFile  = "data/update_method.tpl"
	deleteMethodTemplateFile  = "data/delete_method.tpl"
	typeTemplateFile          = "data/type.tpl"
	dataNewTemplateFile       = "data/data_new.tpl"
	dataGenTemplateFile       = "data/data_gen.tpl"
)

// Clean deletes all template files
func Clean() error {
	return filex.Clean()
}
