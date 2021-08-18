package web

type postRequest struct {
	Title    string `form:"title"`
	Content  string `binding:"required" form:"content"`
	Tags     string `form:"tags"`
	IsPinned bool
}
