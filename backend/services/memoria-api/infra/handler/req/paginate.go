package req

type Paginate struct {
	Page    *int `form:"page"`
	PerPage *int `form:"per_page"`
}
