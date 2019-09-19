package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("c")
	expected := "ccccc"

	if repeated != expected {
		t.Errorf("expected '%s', got '%s'", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
