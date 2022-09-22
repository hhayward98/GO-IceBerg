


type Data struct {
	Dname string

}




func AppRoutes() {

	data := Data{
		Dname: "test"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
		tpl.ExecuteTemplate(w, "index.html", data)

	})

	http.HandleFunc("/Page2", func(w, http.ResponseWriter, r *http.Request) {

		tpl.ExecuteTemplate(w, "Page2.html", data)
	}
}