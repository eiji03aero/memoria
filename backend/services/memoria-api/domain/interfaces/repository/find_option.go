package repository

type FindOptionFilter struct {
	Query string
	Value any
}

type FindOptionJoin struct {
	Query string
}

type FindOption struct {
	Filters []*FindOptionFilter
	Joins   []*FindOptionJoin
	Order   string
}
