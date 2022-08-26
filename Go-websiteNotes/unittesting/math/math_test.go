package math

import (

	"testing"

)

func Test1(t *testing.T) {

	got := Add(5,10)
	want := 15

	if got != want {
		t.Errorf("got %q, wanted %q", got , want)
	}
}