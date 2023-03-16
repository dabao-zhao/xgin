package template

const (
	FindOneMethod = `
func (rp *{{.lowerStartCamelObject}}Repo) FindOne(ctx context.Context, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) (*biz.{{.upperStartCamelObject}}, error) {
	var data *biz.{{.upperStartCamelObject}}
	res := rp.db.WithContext(ctx).Where("{{.lowerStartCamelPrimaryKey}} = ?", {{.lowerStartCamelPrimaryKey}}).First(&data)
	return data, res.Error
}
`
)
