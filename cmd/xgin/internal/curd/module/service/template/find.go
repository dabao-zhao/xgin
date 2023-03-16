package template

const (
	FindOneMethod = `
func (s *{{.upperStartCamelObject}}Service) FindOne(ctx context.Context, req *types.FindOne{{.upperStartCamelObject}}Request) (*types.FindOne{{.upperStartCamelObject}}Response, error) {
	data, err := s.useCase.FindOne(ctx, req.{{.upperStartCamelPrimaryKey}})
	if err != nil {
		s.logger.Error(err)
		return nil, errors.New("find one err")
	}
	var resp *types.FindOne{{.upperStartCamelObject}}Response
	_ = copier.Copy(resp, data)
	return resp, nil
}
`
)
