package fireload

import "testing"

func Test_Metadata_Set(t *testing.T) {
	md := Metadata{}
	md.Set("key", "value")

	if md["key"] != "value" {
		t.Fatalf("Expected md[key] to equal %q, but got %q", "value", md["key"])
	}
}

func Test_Metadata_Get(t *testing.T) {
	md := Metadata{"key": "value"}

	if got, ok := md.Get("key"); !ok || got != "value" {
		t.Fatalf("Expected md[key] to equal %q, but got %q", "value", got)
	}
}
