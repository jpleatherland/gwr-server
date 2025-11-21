package server

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world")
}

