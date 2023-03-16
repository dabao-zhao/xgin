package template

const (
	// DeleteMethod defines a delete template
	DeleteMethod = `
func (s *{{.upperStartCamelObject}}Service) Delete(ctx context.Context, req *types.Delete{{.upperStartCamelObject}}Request) (*types.Delete{{.upperStartCamelObject}}Response, error) {
	err := s.useCase.Delete(ctx, req.{{.upperStartCamelPrimaryKey}})
	if err != nil {
		s.logger.Error(err)
		return nil, errors.New("delete err")
	}
	return &types.Delete{{.upperStartCamelObject}}Response{}, nil
}
`
)
