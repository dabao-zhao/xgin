package template

const (
	UpdateMethod = `
func (s *{{.upperStartCamelObject}}Service) Update(ctx context.Context, req *types.Update{{.upperStartCamelObject}}Request) (*types.Update{{.upperStartCamelObject}}Response, error) {
	var data biz.{{.upperStartCamelObject}}
	_ = copier.Copy(&data, req)
	update, err := s.useCase.Update(ctx, &data)
	if err != nil {
		s.logger.Error(err)
		return nil, errors.New("update err")
	}
	var resp *types.Update{{.upperStartCamelObject}}Response
	_ = copier.Copy(resp, update)
	return resp, nil
}
`
)
