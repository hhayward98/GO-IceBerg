package LoginC


func TestLogin(t *testing.T) {

	got := UserInput("Bob", "password")
	want := "Login Success!"

	if got != want {
		t.Errorf("got %q, wanted %q", got , want)
	}
}


func BenchmarkLogin(b *testing.B) {
	for i :=0; i < b.N ; i++ {
		UserInput("Bob", "password")
	}
}