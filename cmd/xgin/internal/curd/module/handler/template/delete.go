package template

const (
	// DeleteMethod defines a delete template
	DeleteMethod = `
func (h *{{.upperStartCamelObject}}Handler) Delete(ctx context.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		del, err := h.s.Delete(ctx, &types.Delete{{.upperStartCamelObject}}Request{})
		if err != nil {
			h.logger.Error(err)
			ctx.JSON(200, "delete err")
			return
		}
		var resp types.Delete{{.upperStartCamelObject}}Response
		_ = copier.Copy(&resp, del)
		ctx.JSON(200, resp)
	}
}
`
)
