package RegisterC

import (
	"testing"
)


// Test results
func TestRegisterUser(t *testing.T) {

	got := RegisterUser("Bob", "Bob@outlook.com", "password", "password")
	want := "Valid"

	if got != want {
		t.Errorf("got %q, wanted %q", got , want)
	}
}

func TestInputHandler(t *testing.T) {

	got := InputHandler("Bob", "Bob@outlook.com", "password", "password")
	want := "Valid"

	if got != want {
		t.Errorf("got %q, wanted %q", got , want)
	}
}


func TestQueryHandler1(t *testing.T) {
	got := QueryHandler("tom'WHERE 1 = 1; DELETE *TABLES;'")
	want := false

	if got != want {
		t.Errorf("got %t, wanted %t", got , want)
	}
}

// Test for input handling 

// Email is not valid
func TestIPHRfail(t *testing.T) {
	got := InputHandler("Bob", "foo-bar", "password", "password")
	want := "Email is not valid!"
	if got != want {
		t.Errorf("got %q, wanted %q", got , want)
	}
}



func BenchmarkRegisterUser(b *testing.B) {
	for i :=0; i < b.N ; i++ {
		RegisterUser("Bob", "Bob@outlook.com", "password", "password")
	}
}

// func BenchmarkdbUnameCheck(b *testing.B) {
// 	for i :=0; i < b.N ; i++ {
// 		dbUnameCheck("Bob", "Bob@outlook.com")
// 	}
// }

// func BenchmarkInputHandler(b *testing.B) {
// 	for i :=0; i < b.N ; i++ {
// 		InputHandler("Bob", "Bob@outlook.com", "password", "password")
// 	}
// }

// func BenchmarkQueryHandler(b *testing.B) {
// 	for i :=0; i < b.N ; i++ {
// 		QueryHandler("tom'WHERE 1 = 1; DELETE *TABLES;'")
// 	}
// }