package repository

type FindOptionFilter struct {
	Query string
	Value any
}

type FindOptionJoin struct {
	Query string
}

type FindOption struct {
	Filters  []*FindOptionFilter
	Filter   map[string]any
	Joins    []*FindOptionJoin
	Order    string
	Offset   *int
	Limit    *int
	Cursor   string
	CBefore  int
	CAfter   int
	CExclude bool
	Preloads []string
}

func (f *FindOption) Merge(fOpt *FindOption) {
	for _, pl := range fOpt.Preloads {
		f.Preloads = append(f.Preloads, pl)
	}
}
