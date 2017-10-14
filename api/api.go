package main

import (
	"net/http"
	"github.com/AVenezuela/widgets-spa/api/route"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()

	r.GET("/users", route.GetUsers)
	r.GET("/users/:id", route.GetUser)
	r.GET("/widgets/", route.GetWidgets)
	r.GET("/widgets/:id", route.GetWidget)
	r.POST("/widgets/", route.InsertWidget)
	r.PUT("/widgets/", route.UpdateWidget)

	r.GET("/errors/:id", route.GetErrors)

	panic(http.ListenAndServe(":666", r))
}
