package req

type Paginate struct {
	Page    *int `form:"page"`
	PerPage *int `form:"per_page"`
}

type CPaginate struct {
	Cursor   *string `form:"cursor"`
	CBefore  *int    `form:"cbefore"`
	CAfter   *int    `form:"cafter"`
	CExclude *string `form:"cexclude"`
}
