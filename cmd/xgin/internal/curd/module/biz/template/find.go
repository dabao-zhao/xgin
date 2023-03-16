package template

const (
	FindOneMethod = `
func (uc *{{.upperStartCamelObject}}UseCase) FindOne(ctx context.Context, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) (*{{.upperStartCamelObject}}, error) {
	return uc.repo.FindOne(ctx, {{.lowerStartCamelPrimaryKey}})
}
`
	// FindOneInterface defines find row method.
	FindOneInterface = `FindOne(ctx context.Context, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) (*{{.upperStartCamelObject}}, error)`
)
