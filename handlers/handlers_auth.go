package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

var CookieStore *sessions.CookieStore

type BaseAuthHandler struct {
}

func NewBaseAuthHandler(host string) *BaseAuthHandler {
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
