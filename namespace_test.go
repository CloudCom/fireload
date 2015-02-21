package fireload

import "testing"

func Test_NewNamespace(t *testing.T) {
	domain := "test-domain"
	ns := NewNamespace(domain)

	if ns.Domain != domain {
		t.Fatalf("Expected domain to be %q, but got %q", domain, ns.Domain)
	}
}

func Test_Namespace_String(t *testing.T) {
	domain := "test-domain"
	ns := NewNamespace(domain)

	if ns.String() != domain {
		t.Fatalf("Expected Namespace.String() to return %q, but got %q", domain, ns.Domain)
	}
}
