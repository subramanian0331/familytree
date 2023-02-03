package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/subramanian0331/familytree/handlers"
	"github.com/subramanian0331/familytree/store"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize(user, password, dbname string, host string) {
	a.Router = mux.NewRouter()
	gStorage := store.NewStorage(host)
	uStorage := store.NewStorage(host)
	h := handlers.NewBaseHandler(host)
	ah := handlers.NewBaseAuthHandler(host)
	a.initializeRoutes(h, ah)
	// a.initializeRoutes(h *handlers.BaseHandler)
}

func (a *App) initializeRoutes(h *handlers.BaseHandler, ah *handlers.BaseAuthHandler) {
	a.Router.HandleFunc("/health/", h.HealthCheck).Methods("GET")
	a.Router.HandleFunc("/addMember/", h.AddMember).Methods("POST")
	a.Router.HandleFunc("/getMember/{memberId}", h.GetMember).Methods("GET")
	a.Router.HandleFunc("/", h.GetIndex).Methods("GET")
	a.Router.HandleFunc("/auth/{provider}/callback", ah.AuthCallBack).Methods("GET")
	a.Router.HandleFunc("/auth/{provider}", ah.AuthHandle).Methods("GET")
	// a.Router.HandleFunc("/GetMember", h.GetMember).Methods("POST")
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":"+addr, a.Router))
}

func main() {
	fmt.Println("Welcome to Family Tree")
	port := flag.String("listening port", "3333", "provide the listening port for the application. eg \"3333\"")
	flag.Parse()

	app := App{}
	app.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
		"127.0.0.1:6379")

	app.Run(*port)
}
