package route

import (
	"net/http"	
	"github.com/AVenezuela/widgets-spa/api/sqlconn"

	"github.com/gin-gonic/gin"
)

//GetUsers gets all users
func GetUsers(c *gin.Context) {		
	if db, err := sqlconn.NewWidgetDB(nil); err == nil{
		defer db.Close()
		if rows, err := db.All("SELECT [id],[name],[gravatar] FROM [dbo].[User] ORDER BY [Name] ASC"); err == nil{
			c.JSON(http.StatusOK, rows)				
		}else{
			c.JSON(http.StatusBadRequest, err)
		}
	}else{
		c.JSON(http.StatusBadRequest, err)
	}
}

//GetUser gets the user by Id
func GetUser(c *gin.Context) {	
	id := c.Param("id")
	if db, err := sqlconn.NewWidgetDB(nil); err == nil{
		defer db.Close()
		if row, err := db.One("SELECT [id],[name],[gravatar] FROM [dbo].[User] WHERE Id = ?", id); err == nil{
			c.JSON(http.StatusOK, row)
		}else{
			c.JSON(http.StatusBadRequest, err)
		}
	}else{
		c.JSON(http.StatusBadRequest, err)
	}
}
