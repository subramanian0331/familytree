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

type BaseAuthHandler struct {
	store *sessions.CookieStore
}

func NewBaseAuthHandler(host string) *BaseAuthHandler {
	key := "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"
	maxAge := 86400 * 30 // 30 days
	isProd := true       // Set to true when serving over https
	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = isProd
	gothic.Store = store
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
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	t, _ := template.ParseFiles("ui/templates/success.html")
	t.Execute(w, user)

}
