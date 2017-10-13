package route

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"widgets-spa/api/sqlconn"

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

//GetUsers gets all users
func GetWidgets(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	db, err := sqlconn.NewWidgetDB(nil)
	defer db.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rows, err := db.All("SELECT [id],[name],[color],[price],[inventory],[melts] FROM [dbo].[Widget] ORDER BY [Name] ASC")
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
func GetWidget(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	db, err := sqlconn.NewWidgetDB(nil)
	defer db.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	row, err := db.One("SELECT [id],[name],[color],[price],[inventory],[melts] FROM [dbo].[Widget] WHERE Id = ?", id)
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

func InsertWidget(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var widget Widget
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &widget); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	db, err := sqlconn.NewWidgetDB(nil)
	defer db.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	mapWidget := structs.Map(widget)
	delete(mapWidget, "Id")
	id, err := db.Insert("Widget", mapWidget)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	widget.Id = id

	db.Close()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	js, err := json.Marshal(widget)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}

func UpdateWidget(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	var widget Widget
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &widget); err != nil {
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	db, err := sqlconn.NewWidgetDB(nil)
	defer db.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	mapWidget := structs.Map(widget)
	delete(mapWidget, "Id")

	recordsAffected, err := db.Update("Widget", mapWidget, "Id = ?", widget.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if recordsAffected == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	db.Close()
	w.WriteHeader(http.StatusCreated)
	js, err := json.Marshal(widget)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}
