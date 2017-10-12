package users

import (
	//"github.com/AVenezuela/widgets-spa/api/sqlconn"
	"net/http"
	"encoding/json"
)

//User class for user
//Id unique identifier
type User struct {
	ID int
	Name string
	Gravatar string
}

//getUsers gets all users
func GetUsers(w http.ResponseWriter, r *http.Request){

	user := User{1, "Venê", ""}

	js, err := json.Marshal(user)
	if err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}