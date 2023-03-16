package template

const (
	CreateMethod = `
func (uc *{{.upperStartCamelObject}}UseCase) Create(ctx context.Context, data *{{.upperStartCamelObject}}) (*{{.upperStartCamelObject}}, error) {
	return uc.repo.Create(ctx, data)
}
`

	// CreateMethodInterface defines an interface method template for Create code in model
	CreateMethodInterface = `Create(ctx context.Context, data *{{.upperStartCamelObject}}) (*{{.upperStartCamelObject}},error)`
)
