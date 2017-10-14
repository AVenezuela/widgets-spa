package route

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"github.com/AVenezuela/widgets-spa/api/sqlconn"

	"github.com/fatih/structs"
	"github.com/julienschmidt/httprouter"
)

type Widget struct {
	Id        int64  `json:"id" structs:",omitempty"`
	Name      string `json:"name" structs:",omitempty"`
	Color     string `json:"color" structs:",omitempty"`
	Price     string `json:"price" structs:",omitempty"`
	Inventory int    `json:"inventory" structs:",omitempty"`
	Melts     bool   `json:"melts" structs:",omitempty"`
}

//GetWidgets gets all wdigets
func GetWidgets(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	if db, err := sqlconn.NewWidgetDB(nil); err == nil{
		defer db.Close()
		if rows, err := db.All("SELECT [id],[name],[color],[price],[inventory],[melts] FROM [dbo].[Widget] ORDER BY [Name] ASC"); err == nil {
			if js, err := json.Marshal(rows); err == nil{
				w.Write(js)
				return
			}
		}
	}
	w.WriteHeader(http.StatusBadRequest)	
	js := GetJSONError("Failed to get widgets")
	w.Write(js)
}

//GetWidget gets the widget by Id
func GetWidget(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	id := p.ByName("id")
	if db, err := sqlconn.NewWidgetDB(nil); err == nil{
		defer db.Close()
		if row, err := db.One("SELECT [id],[name],[color],[price],[inventory],[melts] FROM [dbo].[Widget] WHERE Id = ?", id); err == nil {
			if js, err := json.Marshal(row); err == nil{
				w.Write(js)
				return
			}
		}
	}
	w.WriteHeader(http.StatusNotFound)
	js := GetJSONError("Failed to get widget")
	w.Write(js)
}

//InsertWidget inserts new widget
func InsertWidget(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")	
	if body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)); err == nil {
		var widget Widget
		if err := r.Body.Close(); err != nil {
			panic(err)
		}	
		if err := json.Unmarshal(body, &widget); err != nil {			
			w.WriteHeader(http.StatusUnprocessableEntity)
			if err := json.NewEncoder(w).Encode(err); err != nil {
				panic(err)
			}
		}
		if db, err := sqlconn.NewWidgetDB(nil); err == nil{
			defer db.Close()
			mapWidget := structs.Map(widget)
			delete(mapWidget, "Id")
			if id, err := db.Insert("Widget", mapWidget); err == nil{				
				if js, err := json.Marshal(widget); err == nil{
					widget.Id = id
					w.WriteHeader(http.StatusCreated)
					w.Write(js)
					return
				}
			}
		}
	}
	w.WriteHeader(http.StatusUnprocessableEntity)
	js := GetJSONError("Failed to insert new widget")
	w.Write(js)
}

func UpdateWidget(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var widget Widget
	if body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)); err == nil {
		if err := r.Body.Close(); err != nil {
			panic(err)
		}
		if err := json.Unmarshal(body, &widget); err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			if err := json.NewEncoder(w).Encode(err); err != nil {
				panic(err)
			}
		}
		if db, err := sqlconn.NewWidgetDB(nil); err == nil{
			defer db.Close()
			mapWidget := structs.Map(widget)
			delete(mapWidget, "Id")	
			if recordsAffected, err := db.Update("Widget", mapWidget, "Id = ?", widget.Id); err == nil{
				if recordsAffected != 0 {					
					w.WriteHeader(http.StatusOK)
					if js, err := json.Marshal(widget); err == nil {
						w.Write(js)
						return
					}				
				}
			}
		}
	}	
	w.WriteHeader(http.StatusNotFound)
	js := GetJSONError("Failed to update widget")
	w.Write(js)
}
