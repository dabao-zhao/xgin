package template

// NewTpl defines the template for creating model instance.
const NewTpl = `
func New{{.upperStartCamelObject}}Service(useCase *biz.{{.upperStartCamelObject}}UseCase, logger *logging.Logger) *{{.upperStartCamelObject}}Service {
	return &{{.upperStartCamelObject}}Service{
		useCase : useCase,
		logger  : logger,
	}
}
`
