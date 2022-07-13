file, err := os.Open('/path/to/file')
if err != nil {
	panic(err)
}
defer file.Close()

buf := make([]byte, BUFSIZE)
for {
	n, err := file.Read(buf)
	if n == 0{
		break
	}
	if err != nil {
		panic (err)
	}

	fmt.Print(string(buf[:n]))
}



































