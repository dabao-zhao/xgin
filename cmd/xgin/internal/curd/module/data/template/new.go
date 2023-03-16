package template

// NewTpl defines the template for creating model instance.
const NewTpl = `
func New{{.upperStartCamelObject}}Repo(db *gorm.DB, logger *logging.Logger) biz.{{.upperStartCamelObject}}Repo {
	return &{{.lowerStartCamelObject}}Repo{
		db 	   : db,
		logger : logger,
	}
}
`
