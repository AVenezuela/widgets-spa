package main

import (	
	"net/http"
	"github.com/AVenezuela/widgets-spa/api/routes"
)

func main() {
	http.HandleFunc("/users", users.GetUsers)

	panic(http.ListenAndServe(":666", nil))
}