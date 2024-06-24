package pagination

type Page[T any] struct {
	data   []T
	total  int
	limit  int
	offset int
}

func NewPage[T any](data []T, total int, limit int, offset int) Page[T] {
	return Page[T]{data: data, total: total, limit: limit, offset: offset}
}

func NewEmptyPage[T any]() Page[T] {
	return Page[T]{data: make([]T, 0), total: 0, limit: 0, offset: 0}
}

func (p Page[T]) GetObjects() []T {
	return p.data
}

func (p Page[T]) GetTotal() int {
	return p.total
}

func (p Page[T]) GetLimit() int {
	return p.limit
}

func (p Page[T]) GetOffset() int {
	return p.offset
}
