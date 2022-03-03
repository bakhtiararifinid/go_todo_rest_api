package todo

type Request struct {
	Title string `json:"title" binding:"required"`
}
