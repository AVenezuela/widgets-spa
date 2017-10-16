package main

import (
	"net/http"
	"time"

	"github.com/AVenezuela/widgets-spa/api/route"
	"github.com/gin-contrib/cors"

	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	authMiddleware := &jwt.GinJWTMiddleware{
		Realm:      "test zone",
		Key:        []byte("secret key"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(userId string, password string, c *gin.Context) (string, bool) {
			if (userId == "admin" && password == "admin") || (userId == "test" && password == "test") {
				return userId, true
			}

			return userId, false
		},
		Authorizator: func(userId string, c *gin.Context) bool {
			if userId == "admin" {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header:Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	}

	router.POST("/login", authMiddleware.LoginHandler)

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:555"}
	config.AllowHeaders = []string{"Authorization"}
	config.AllowMethods = []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"}
	router.Use(cors.New(config))

	auth := router.Group("/api")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		/*WIDGETS ROUTES */
		auth.GET("/widgets", route.GetWidgets)
		auth.POST("/widget", route.InsertWidget)
		auth.PUT("/widget", route.UpdateWidget)
		auth.DELETE("/widget/:id", route.DeleteWidget)
		auth.GET("/widget/:id", route.GetWidget)

		/*USERS ROUTES */
		auth.GET("/users", route.GetUsers)
		auth.POST("/user", route.InsertUser)
		auth.PUT("/user", route.UpdateUser)
		auth.DELETE("/user/:id", route.DeleteUser)
		auth.GET("/user/:id", route.GetUser)
	}

	http.ListenAndServe(":666", router)
}
