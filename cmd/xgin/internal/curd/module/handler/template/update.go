package template

const (
	UpdateMethod = `
func (h *{{.upperStartCamelObject}}Handler) Update(ctx context.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		update, err := h.s.Update(ctx, &types.{{.upperStartCamelObject}}Request{})
		if err != nil {
			h.logger.Error(err)
			ctx.JSON(200, "update err")
			return
		}
		var resp types.Update{{.upperStartCamelObject}}Response
		_ = copier.Copy(&resp, update)
		ctx.JSON(200, resp)
	}
}
`
)
