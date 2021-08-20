package web

type siteparameterRequest struct {
	Name        string `binding:"required" form:"name"`
	Domain      string `binding:"required" form:"domain"`
	Description string `binding:"required" form:"description"`
	PageSize    int    `binding:"required" form:"page_size"`
	Separator   string `binding:"required" form:"separator"`
}
