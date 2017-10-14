package route

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/AVenezuela/widgets-spa/api/sqlconn"

	"github.com/julienschmidt/httprouter"
)

func GetErrors(w http.ResponseWriter, r *http.Request, p httprouter.Params){	
	var err error
	erroId := p.ByName("id")
	w.Header().Set("Content-Type", "application/json")
	if i, err := strconv.ParseInt(erroId, 10, 32); err == nil{
		w.WriteHeader(int(i))
		return
	}
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

//GetUsers gets all users
func GetUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {	
	w.Header().Set("Content-Type", "application/json")	
	if db, err := sqlconn.NewWidgetDB(nil); err == nil{
		defer db.Close()
		if rows, err := db.All("SELECT [id],[name],[gravatar] FROM [dbo].[User] ORDER BY [Name] ASC"); err == nil{
			if js, err := json.Marshal(rows); err == nil{					
				w.Write(js)
				return
			}				
		}				
	}
	w.WriteHeader(http.StatusBadRequest)	
	js := GetJSONError("Failed to get users")
	w.Write(js)
}

//GetUser gets the user by Id
func GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	id := p.ByName("id")
	if db, err := sqlconn.NewWidgetDB(nil); err == nil{
		defer db.Close()
		if row, err := db.One("SELECT [id],[name],[gravatar] FROM [dbo].[User] WHERE Id = ?", id); err == nil{
			if js, err := json.Marshal(row); err == nil {
				w.Write(js)					
				return
			}						
		}
	}
	w.WriteHeader(http.StatusNotFound)
	js := GetJSONError("Failed to get user")
	w.Write(js)
}
