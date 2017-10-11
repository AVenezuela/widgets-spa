package main

import "net/http"

func main() {
	http.HandleFunc("/", homeHandler)	
	serveFolder("/components/")
	serveFolder("/img/")
	serveFolder("/scripts/")

	panic(http.ListenAndServe(":17901", nil))
}

func serveFolder(folderName string){
	http.HandleFunc(folderName, serveFile)
}

func serveFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
