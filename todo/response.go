package todo

type Response struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	IsCompleted bool   `json:"is_completed"`
}
