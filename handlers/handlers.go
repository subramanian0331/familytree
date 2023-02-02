package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/subramanian0331/familytree/models"
	"github.com/subramanian0331/familytree/store"
)

type BaseHandler struct {
	DB store.Storage
}

func NewBaseHandler(host string) *BaseHandler {
	bh := BaseHandler{
		DB: store.NewStorage(host),
	}
	return &bh
}

func (b *BaseHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Health OK 100%")
}

func (b *BaseHandler) AddMember(w http.ResponseWriter, r *http.Request) {
	m := models.Member{
		Id: uuid.New(),
	}
	// // Buffer the body
	// if r.Body != nil {
	// 	bodyBytes, err := ioutil.ReadAll(r.Body)
	// 	if err != nil {
	// 		fmt.Printf("Body reading error: %v", err)
	// 		return
	// 	}
	// 	defer r.Body.Close()
	// 	if len(bodyBytes) > 0 {
	// 		var prettyJSON bytes.Buffer
	// 		if err = json.Indent(&prettyJSON, bodyBytes, "", "\t"); err != nil {
	// 			fmt.Printf("JSON parse error: %v", err)
	// 			return
	// 		}
	// 		fmt.Println(string(prettyJSON.Bytes()))
	// 	} else {
	// 		fmt.Printf("Body: No Body Supplied\n")
	// 	}
	// }

	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		fmt.Println("error", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("herere: %+v", m)
	err = b.DB.AddMember(&m)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Failed to Add Member to DB", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Person Updated to DB")
}

func (b *BaseHandler) GetMember(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	fmt.Println(params)

}
