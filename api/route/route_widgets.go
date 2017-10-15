package route

import (		
	"net/http"
	"github.com/AVenezuela/widgets-spa/api/sqlconn"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

//Widget interface
type Widget struct {
	Id        int64  `json:"id" structs:",omitempty"`
	Name      string `json:"name" structs:",omitempty"`
	Color     string `json:"color" structs:",omitempty"`
	Price     string `json:"price" structs:",omitempty"`
	Inventory int    `json:"inventory" structs:",omitempty"`
	Melts     bool   `json:"melts"`
}

//GetWidgets gets all wdigets
func GetWidgets(c *gin.Context) {	
	if db, err := sqlconn.NewWidgetDB(nil); err == nil{
		defer db.Close()
		if rows, err := db.All("SELECT [id],[name],[color],[price],[inventory],[melts] FROM [dbo].[Widget] ORDER BY [Name] ASC"); err == nil {
			c.JSON(http.StatusOK, rows)
		}else{
			c.JSON(http.StatusBadRequest, err)
		}
	}else{
		c.JSON(http.StatusBadRequest, err)
	}
}

//GetWidget gets the widget by Id
func GetWidget(c *gin.Context) {	
	id := c.Param("id")
	if db, err := sqlconn.NewWidgetDB(nil); err == nil{
		defer db.Close()
		if row, err := db.One("SELECT [id],[name],[color],[price],[inventory],[melts] FROM [dbo].[Widget] WHERE Id = ?", id); err == nil {
			c.JSON(http.StatusOK, row)
		}else{
			c.JSON(http.StatusBadRequest, err)
		}
	}else{
		c.JSON(http.StatusBadRequest, err)
	}
}

//InsertWidget inserts new widget
func InsertWidget(c *gin.Context) {
	var widget Widget
	if err := c.BindJSON(&widget); err == nil{
		if db, err := sqlconn.NewWidgetDB(nil); err == nil{
			defer db.Close()
			mapWidget := structs.Map(widget)
			delete(mapWidget, "Id")
			if id, err := db.Insert("Widget", mapWidget); err == nil{				
				widget.Id = id
				c.JSON(http.StatusCreated, widget)
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

//UpdateWidget update widget record
func UpdateWidget(c *gin.Context) {
	var widget Widget
	if err := c.BindJSON(&widget); err == nil{
		if db, err := sqlconn.NewWidgetDB(nil); err == nil{
			defer db.Close()
			mapWidget := structs.Map(widget)
			delete(mapWidget, "Id")
			if recordsAffected, err := db.Update("Widget", mapWidget, "Id = ?", widget.Id); err == nil{
				if recordsAffected != 0 {
					c.JSON(http.StatusCreated, widget)
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

//DeleteWidget delete widget record
func DeleteWidget(c *gin.Context) {
	id := c.Param("id")
	if db, err := sqlconn.NewWidgetDB(nil); err == nil{
		defer db.Close()
		if recordsAffected, err := db.Delete("Widget", "Id = ?", id); err == nil{
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
