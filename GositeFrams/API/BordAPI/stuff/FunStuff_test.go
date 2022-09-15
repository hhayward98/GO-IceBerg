package BoardAPI

import (
	"testing"
)


func TestAdd(t *testing.T) {

	got := Add(5,10)
	want := 15

	if got != want {
		t.Errorf("got %q, wanted %q", got , want)
	}
}

func TestSubtract(t *testing.T) {
	got := Subtract(5,10)
	want := 5

	if got != want {
		t.Errorf("got %q, wanted %q", got , want)
	}
}

func BenchmarkAdd(b *testing.B) {
	for i :=0; i < b.N ; i++ {
		Add(4, 6)
	}
}