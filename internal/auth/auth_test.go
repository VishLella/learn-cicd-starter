package auth

import (
	"testing"
)

func TestAPIKey(t *testing.T) {
	got := "random"
	want := "random"
	if got == want {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
