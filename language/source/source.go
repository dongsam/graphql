package source

const defaultName = "GraphQL"

type Source struct {
	Name string
	Body string
}

func NewSource(s *Source) *Source {
	if s == nil {
		s = &Source{}
	}

	s.Name = defaultName

	return s
}
