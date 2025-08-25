package auth

import (
	"reflect"
	"testing"
)

func TestAut(t *testing.T) {
	test := map[string]struct {
		name string
		got  string
		want string
	}{
		"simple": {}
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v got: %v", got, want)
	}
}
