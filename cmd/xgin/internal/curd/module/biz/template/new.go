package template

// NewTpl defines the template for creating model instance.
const NewTpl = `
func New{{.upperStartCamelObject}}UseCase(repo {{.upperStartCamelObject}}Repo, logger *logging.Logger) *{{.upperStartCamelObject}}UseCase {
	return &{{.upperStartCamelObject}}UseCase{
		repo  : repo,
		logger: logger,
	}
}
`
