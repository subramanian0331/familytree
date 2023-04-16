package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/subramanian0331/familytree/models"
	"github.com/subramanian0331/familytree/store"
)

type BaseHandler struct {
	familyDB store.Storage
	userDB   store.UserStorage
}

func NewBaseHandler(graphDB store.Storage, userDB store.UserStorage) *BaseHandler {
	bh := BaseHandler{
		familyDB: graphDB,
		userDB:   userDB,
	}
	return &bh
}

// Basic service health check.
func (b *BaseHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Health OK")
}

//Adds new family member to the database
func (b *BaseHandler) AddMember(w http.ResponseWriter, r *http.Request) {
	m := models.Member{
		Id: uuid.New(),
	}

	// un-marshaling json body to member
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		fmt.Println("error", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = b.familyDB.AddMember(&m)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Failed to Add Member to DB", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Person Updated to DB")
}

// Get Family member profile
func (b *BaseHandler) GetMember(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	fmt.Println(params)

}

func (b *BaseHandler) GetIndex(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("ui/templates/index.html")
	t.Execute(w, false)
}
