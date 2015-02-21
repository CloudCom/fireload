package fireload

// Namespace .
type Namespace struct {
	Domain   string
	Metadata Metadata
}

// NewNamespace .
func NewNamespace(domain string) Namespace {
	return Namespace{
		Domain:   domain,
		Metadata: Metadata{},
	}
}

// String .
func (n Namespace) String() string {
	return n.Domain
}
