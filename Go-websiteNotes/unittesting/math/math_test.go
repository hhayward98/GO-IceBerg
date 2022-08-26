package math

import (
	"fmt"
	"testing"

)

// arg1 + arg2 = expected
type addTest struct {
	arg1, arg2, expected int
}

var AddTest = []addTest {
	addTest{5, 10, 15},
	addTest{3, 8, 11},
	addTest{30, 49, 79},

}

func Test1(t *testing.T) {

	got := Add(5,10)
	want := 15

	if got != want {
		t.Errorf("got %q, wanted %q", got , want)
	}
}

func TestAdd(t *testing.T) {
	for _, test := range AddTest {
		if output := Add(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}
// when creating functions for Benchmark testting
// Use "Benchmark" as prefix and use Function Name as suffix  
func BenchmarkAdd(b *testing.B) {
	for i :=0; i < b.N ; i++ {
		Add(4, 6)
	}
}

func ExampleAdd() {
	fmt.Println(Add(8,2))

}

