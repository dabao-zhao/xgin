package template

const (
	UpdateMethod = `
func (uc *{{.upperStartCamelObject}}UseCase) Update(ctx context.Context, data *{{.upperStartCamelObject}}) (*{{.upperStartCamelObject}}, error) {
	return uc.repo.Update(ctx, data)
}
`

	// UpdateMethodInterface defines an interface method template for Update code in model
	UpdateMethodInterface = `Update(ctx context.Context, data *{{.upperStartCamelObject}}) (*{{.upperStartCamelObject}},error)`
)
