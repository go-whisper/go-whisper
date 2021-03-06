package web

type postRequest struct {
	Title    string `form:"title"`
	Content  string `binding:"required" form:"content"`
	URL      string `form:"url"`
	Tags     string `form:"tags"`
	IsPinned bool   `form:"is_pinned"`
}
