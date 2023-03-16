package template

const (
	// DeleteMethod defines a delete template
	DeleteMethod = `
func (uc *{{.upperStartCamelObject}}UseCase) Delete(ctx context.Context, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) error {
	return uc.repo.Delete(ctx, {{.lowerStartCamelPrimaryKey}})
}
`

	// DeleteMethodInterface defines a delete template for interface method
	DeleteMethodInterface = `Delete(ctx context.Context, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) error`
)
