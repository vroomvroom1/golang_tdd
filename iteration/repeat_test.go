package iteration

import "testing"

func TestRepeat(t *testing.T) {
	values := Repeat("c")
	expected := "ccccc"

	if values != expected {
		t.Errorf("expected '%s', got '%s'", expected, values)
	}
}
