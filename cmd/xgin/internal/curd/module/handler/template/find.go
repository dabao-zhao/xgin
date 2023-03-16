package template

const (
	FindOneMethod = `
func (h *{{.upperStartCamelObject}}Handler) FindOne(ctx context.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		one, err := h.s.FindOne(ctx, &types.FindOne{{.upperStartCamelObject}}Request{})
		if err != nil {
			h.logger.Error(err)
			ctx.JSON(200, "find err")
			return
		}
		var resp types.FindOne{{.upperStartCamelObject}}Response
		_ = copier.Copy(&resp, one)
		ctx.JSON(200, resp)
	}
}
`
)
