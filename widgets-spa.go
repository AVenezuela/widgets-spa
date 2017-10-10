package main

import "net/http"

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/components/", serveFile)
	http.HandleFunc("/img/", serveFile)
	http.HandleFunc("/scripts/", serveFile)

	panic(http.ListenAndServe(":17901", nil))
}

func serveFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
