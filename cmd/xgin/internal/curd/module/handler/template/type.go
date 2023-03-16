package template

// TypeTpl defines a template for types in handler.
const TypeTpl = `
type (
	{{.upperStartCamelObject}}Handler struct {
		s       *services.{{.upperStartCamelObject}}Service
		logger  *logging.Logger
	}
)
`
