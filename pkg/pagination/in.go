package pagination

type PaginationContext struct {
	limit  int
	offset int
}

func NewPaginationContext(limit int, offset int) PaginationContext {
	return PaginationContext{limit: limit, offset: offset}
}

func (p PaginationContext) GetLimit() int {
	return p.limit
}

func (p PaginationContext) GetOffset() int {
	return p.offset
}
