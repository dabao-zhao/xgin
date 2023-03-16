package template

const (
	CreateMethod = `
func (s *{{.upperStartCamelObject}}Service) Create(ctx context.Context, req *types.Create{{.upperStartCamelObject}}Request) (*types.Create{{.upperStartCamelObject}}Response, error) {
	var data biz.{{.upperStartCamelObject}}
	_ = copier.Copy(&data, req)
	create, err := s.useCase.Create(ctx, &data)
	if err != nil {
		s.logger.Error(err)
		return nil, errors.New("create err")
	}
	var resp *types.Create{{.upperStartCamelObject}}Response
	_ = copier.Copy(resp, create)
	return resp, nil
}
`
)
