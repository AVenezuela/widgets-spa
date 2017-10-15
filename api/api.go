package main

import (	
	"github.com/AVenezuela/widgets-spa/api/route"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:555"}	
	config.AllowMethods = []string{"PUT", "DELETE", "POST", "GET"}
	router.Use(cors.New(config))
	
	/*WIDGETS ROUTES */
	router.GET("/widgets", route.GetWidgets)
	router.POST("/widget", route.InsertWidget)
	router.PUT("/widget", route.UpdateWidget)	
	router.DELETE("/widget/:id", route.DeleteWidget)	
	router.GET("/widget/:id", route.GetWidget)	
	
	/*USERS ROUTES */
	router.GET("/users", route.GetUsers)
	router.POST("/user", route.InsertUser)
	router.PUT("/user", route.UpdateUser)
	router.DELETE("/user/:id", route.DeleteUser)
	router.GET("/user/:id", route.GetUser)

	
	router.Run(":666")

	//handler := cors.Default().Handler(router)
	//r.GET("/errors/:id", route.GetErrors)
	//panic(http.ListenAndServe(":666", handler))
}
