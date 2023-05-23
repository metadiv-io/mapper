package mapper

type BaseMapper[T any] struct {
	BeforeMap2Model func(from any) any
	AfterMap2Model  func(from any, to *T) *T
}

func (m *BaseMapper[T]) Map2Model(from any) *T {
	from = neverBePtr(from)
	if m.BeforeMap2Model != nil {
		from = m.BeforeMap2Model(from)
	}
	to := Map2Model[T](from)
	if m.AfterMap2Model != nil {
		to = m.AfterMap2Model(from, to)
	}
	return to
}

func (m *BaseMapper[T]) Map2Models(from []any) []T {
	var to []T
	for _, f := range from {
		to = append(to, *m.Map2Model(f))
	}
	return to
}
