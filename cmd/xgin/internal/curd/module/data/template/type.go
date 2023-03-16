package template

// TypeTpl defines a template for types in service.
const TypeTpl = `
var _ biz.{{.upperStartCamelObject}}Repo = (*{{.lowerStartCamelObject}}Repo)(nil)

type (
	{{.lowerStartCamelObject}}Repo struct {
		db     *gorm.DB
		logger *logging.Logger
	}
)
`
