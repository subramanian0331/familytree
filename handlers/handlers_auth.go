package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"github.com/subramanian0331/familytree/models"
	"github.com/subramanian0331/familytree/store"
	"golang.org/x/crypto/bcrypt"
)

var CookieStore *sessions.CookieStore

type BaseAuthHandler struct {
	UserService store.UserStorage
}

func NewBaseAuthHandler(host string, UserService store.UserStorage) *BaseAuthHandler {
	key := "my-secret"
	maxAge := 3600 // 30 days
	CookieStore = sessions.NewCookieStore([]byte(key))
	CookieStore.MaxAge(maxAge)
	CookieStore.Options.Path = "/"
	CookieStore.Options.HttpOnly = true // HttpOnly should always be enabled
	gothic.Store = CookieStore
	goth.UseProviders(google.New(
		"49422998462-fomnp86pputufo8tpaqcr1rdhq43f8pe.apps.googleusercontent.com",
		"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",
		"http://localhost:3333/auth/google/callback",
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile"))
	bh := BaseAuthHandler{}
	bh.UserService = UserService

	return &bh

}

func (h *BaseAuthHandler) AuthHandle(w http.ResponseWriter, r *http.Request) {
	gothic.BeginAuthHandler(w, r)
}

func (h *BaseAuthHandler) AuthCallBack(w http.ResponseWriter, r *http.Request) {

	// Get the session from the store
	session, _ := CookieStore.Get(r, "session-name")

	// Complete the OAuth2 flow and get the user's credentials
	creds, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["authenticated"] = true
	err = session.Save(r, w)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	fmt.Fprintf(w, "User information: %+v", creds)
	t, _ := template.ParseFiles("ui/templates/home.html")
	t.Execute(w, creds)
}

func (b *BaseAuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse the form data", http.StatusBadGateway)
		return
	}

	firstName := r.Form.Get("first_name")
	lastName := r.Form.Get("last_name")
	nickName := r.Form.Get("nick_name")
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	passwordConfirmation := r.Form.Get("password_confirmation")
	gender, err := models.StringToSex(strings.Title(r.Form.Get("gender")))

	fmt.Println("data:", firstName, lastName, email, password, passwordConfirmation)
	// Validate the form data
	if firstName == "" || lastName == "" || email == "" || password == "" || passwordConfirmation == "" || r.Form.Get("gender") == "" {
		http.Error(w, "Please fill in all required fields", http.StatusBadRequest)
		return
	}
	if password != passwordConfirmation {
		http.Error(w, "Passwords do not match", http.StatusBadRequest)
		return
	}
	user := models.User{}
	user.Firstname = firstName
	user.Lastname = lastName
	user.Nickname = nickName
	user.Email = email
	user.UserMetaData.Gender = gender

	passwd := r.Form.Get("password")
	// Hash the password, bcrypt handles the salt internally.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to Generate Hash from password", http.StatusBadGateway)
		return
	}
	user.PassHash = string(hashedPassword)

	err = b.UserService.AddUser(user)
	if err != nil {
		if err.Error() == "user exists" {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		http.Error(w, "failed to create new user."+"error:"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User Create successfully")
}

func (b *BaseAuthHandler) Deactivate(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}

	err = b.UserService.DeleteUser(userId)
	if err != nil {
		http.Error(w, "failed"+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintf(w, "User Deleted successfully")
}

func (b *BaseAuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to Parse form", http.StatusBadRequest)
	}
	// userEmail := r.Form.Get("email")
	// userPassword := r.Form.Get("password")

}
