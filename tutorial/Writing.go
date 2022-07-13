f, err := os.Create("/path/to/file")
if err != nil {
	panic(err)
}
defer f.Close()

b := []byte("Foo")
n, err := f.Write(b)
if err != nil {
	panic(err)
}
fmt.Println(n)




































