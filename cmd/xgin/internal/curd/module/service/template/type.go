package template

// TypeTpl defines a template for types in service.
const TypeTpl = `
type (
	{{.upperStartCamelObject}}Service struct {
		useCase *biz.{{.upperStartCamelObject}}UseCase
		logger  *logging.Logger
	}
)
`
