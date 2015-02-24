package fireload

// Namespace represents a Firebase Namespace (e.g. `<namespace>.firebaseio.com`).
type Namespace struct {
	// Domain is the fully-qualified domain name of the Firebase Namespace.
	// The Domain for the samplechat Namespace mentioned in the Firebase documentation
	// would be `samplechat.firebaseio-demo.com`.
	Domain string

	// Metadata holds any key-value pairs associated with this Namespace; e.g. the admin secret.
	Metadata Metadata
}

// NewNamespace returns a Namespace with the given domain.
func NewNamespace(domain string) Namespace {
	return Namespace{
		Domain:   domain,
		Metadata: Metadata{},
	}
}

// String satisfies the fmt.Stringer interface and returns the Namespace's Domain.
func (n Namespace) String() string {
	return n.Domain
}
