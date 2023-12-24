package filter

type Specification[T any] interface {
	IsSatisfiedBy(item T) bool
}

type AndSpecification[T any] struct {
	Specifications []Specification[T]
}

func AndSpecificationOf[T any](specifications ...Specification[T]) *AndSpecification[T] {
	return &AndSpecification[T]{Specifications: specifications}
}

func (s *AndSpecification[T]) IsSatisfiedBy(item T) bool {
	for _, specification := range s.Specifications {
		if !specification.IsSatisfiedBy(item) {
			return false
		}
	}
	return true
}
