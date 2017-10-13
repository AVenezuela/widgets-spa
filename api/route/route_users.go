package route

import (
	"encoding/json"
	"net/http"
	"widgets-spa/api/sqlconn"

	"github.com/julienschmidt/httprouter"
)

//GetUsers gets all users
func GetUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	db, err := sqlconn.NewWidgetDB(nil)
	defer db.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rows, err := db.All("SELECT [id],[name],[gravatar] FROM [dbo].[User] ORDER BY [Name] ASC")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(rows)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	db.Close()
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

//GetUser gets the user by Id
func GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	db, err := sqlconn.NewWidgetDB(nil)
	defer db.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	row, err := db.One("SELECT [id],[name],[gravatar] FROM [dbo].[User] WHERE Id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(row)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	db.Close()
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
