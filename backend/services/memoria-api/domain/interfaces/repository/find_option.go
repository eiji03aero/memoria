package repository

type FindOptionFilter struct {
	Query string
	Value any
}

type FindOption struct {
	Filters []*FindOptionFilter
}
