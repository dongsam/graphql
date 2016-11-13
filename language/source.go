package language

const defaultName = "GraphQL"

// Source ...
type Source struct {
	Name string
	Body string
}

// NewSource ...
func NewSource(s *Source) *Source {
	if s == nil {
		s = &Source{}
	}

	s.Name = defaultName

	return s
}
