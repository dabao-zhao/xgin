package template

// NewTpl defines the template for creating model instance.
const NewTpl = `
func New{{.upperStartCamelObject}}Handler(useCase *services.{{.upperStartCamelObject}}Service, logger *logging.Logger) *{{.upperStartCamelObject}}Handler {
	return &{{.upperStartCamelObject}}Handler{
		s : s,
		logger  : logger,
	}
}
`
