package main

import (	
	"github.com/AVenezuela/widgets-spa/api/route"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	
	router.POST("/widget", route.InsertWidget)
	router.PUT("/widget", route.UpdateWidget)
	router.GET("/users", route.GetUsers)
	router.GET("/users/:id", route.GetUser)
	router.GET("/widgets", route.GetWidgets)
	router.GET("/widgets/:id", route.GetWidget)	
	
	
	router.Run(":666")

	//handler := cors.Default().Handler(router)
	//r.GET("/errors/:id", route.GetErrors)
	//panic(http.ListenAndServe(":666", handler))
}
