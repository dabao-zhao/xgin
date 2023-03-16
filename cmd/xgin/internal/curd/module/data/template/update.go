package template

const (
	UpdateMethod = `
func (rp *{{.lowerStartCamelObject}}Repo) Update(ctx context.Context, data *biz.{{.upperStartCamelObject}}) (*biz.{{.upperStartCamelObject}}, error) {
	res := rp.db.WithContext(ctx).Updates(&data)
	return data, res.Error
}
`
)
