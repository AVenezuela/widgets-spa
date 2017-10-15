package route

import (
	"net/http"
	"github.com/AVenezuela/widgets-spa/api/sqlconn"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

//User interface
type User struct {
	Id        int64  `json:"id" structs:",omitempty"`
	Name      string `json:"name" structs:",omitempty"`
	Gravatar  string `json:"gravatar" structs:",omitempty"`	
}

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


//InsertUser inserts new User
func InsertUser(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err == nil{
		if db, err := sqlconn.NewWidgetDB(nil); err == nil{
			defer db.Close()
			mapUser := structs.Map(user)
			delete(mapUser, "Id")
			if id, err := db.Insert("[User]", mapUser); err == nil{				
				user.Id = id
				c.JSON(http.StatusCreated, user)
			}else{
				c.JSON(http.StatusBadRequest, err)
			}
		}else{
			c.JSON(http.StatusBadRequest, err)
		}
	}else{
		c.JSON(http.StatusBadRequest, err)
	}
}

//UpdateUser update user record
func UpdateUser(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err == nil{
		if db, err := sqlconn.NewWidgetDB(nil); err == nil{
			defer db.Close()
			mapUser := structs.Map(user)
			delete(mapUser, "Id")
			if recordsAffected, err := db.Update("[User]", mapUser, "Id = ?", user.Id); err == nil{
				if recordsAffected != 0 {
					c.JSON(http.StatusCreated, user)
				}else{
					c.JSON(http.StatusBadRequest, gin.H{
						"message": "No record affected",
					})
				}
			}else{
				c.JSON(http.StatusBadRequest, err)
			}
		}else{
			c.JSON(http.StatusBadRequest, err)
		}
	}else{
		c.JSON(http.StatusBadRequest, err)
	}
}

//DeleteUser delete user record
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if db, err := sqlconn.NewWidgetDB(nil); err == nil{
		defer db.Close()
		if recordsAffected, err := db.Delete("[User]", "Id = ?", id); err == nil{
			if recordsAffected != 0 {
				c.JSON(http.StatusOK, gin.H{
					"message": "Record deleted",
				})
			}else{
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "No record affected",
				})
			}
		}else{
			c.JSON(http.StatusBadRequest, err)
		}
	}else{
		c.JSON(http.StatusBadRequest, err)
	}
}