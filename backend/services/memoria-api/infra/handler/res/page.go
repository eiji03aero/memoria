package res

type Pagination struct {
	CurrentPage int `json:"current_page"`
	TotalPage   int `json:"total_page"`
	PerPage     int `json:"per_page"`
}

type CPagination struct {
	NextCursor string `json:"next_cursor"`
	PrevCursor string `json:"prev_cursor"`
}
