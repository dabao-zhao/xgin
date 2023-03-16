package template

const (
	CreateMethod = `
func (h *{{.upperStartCamelObject}}Handler) Create(ctx context.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		create, err := h.s.Create(ctx, &types.{{.upperStartCamelObject}}Request{})
		if err != nil {
			h.logger.Error(err)
			ctx.JSON(200, "create err")
			return
		}
		var resp types.Create{{.upperStartCamelObject}}Response
		_ = copier.Copy(&resp, create)
		ctx.JSON(200, resp)
	}
}
`
)
