package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/subramanian0331/familytree/handlers"
	"github.com/subramanian0331/familytree/middleware"
	"github.com/subramanian0331/familytree/store"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize(user, password, dbname string, host string, redisPort string, pgPort string) {
	a.Router = mux.NewRouter()
	gStorage := store.NewStorage(host + ":" + redisPort)
	uStorage := store.NewUserStorage(host, pgPort, user, password, dbname)
	h := handlers.NewBaseHandler(gStorage, uStorage)
	ah := handlers.NewBaseAuthHandler(host, uStorage)
	a.initializeRoutes(h, ah)
}

func (a *App) initializeRoutes(h *handlers.BaseHandler, ah *handlers.BaseAuthHandler) {
	a.Router.HandleFunc("/health", middleware.Chain(h.HealthCheck, middleware.Logging(), middleware.SessionValidation(), middleware.Method("GET"))).Methods("GET")
	a.Router.HandleFunc("/addMember/", middleware.Chain(h.AddMember, middleware.Logging(), middleware.Method("POST"), middleware.SessionValidation())).Methods("POST")
	a.Router.HandleFunc("/getMember/{memberId}", middleware.Chain(h.GetMember, middleware.Logging(), middleware.Method("GET"), middleware.SessionValidation())).Methods("GET")
	a.Router.HandleFunc("/", middleware.Chain(h.GetIndex, middleware.Logging(), middleware.Method("GET"))).Methods("GET")
	a.Router.HandleFunc("/auth/{provider}/callback", middleware.Chain(ah.AuthCallBack, middleware.Logging(), middleware.Method("GET"))).Methods("GET")
	a.Router.HandleFunc("/auth/{provider}", ah.AuthHandle).Methods("GET")
	a.Router.HandleFunc("/signup/", middleware.Chain(ah.Signup, middleware.Logging(), middleware.Method("POST"))).Methods("POST")
	a.Router.HandleFunc("/deactivate/user/{id}", middleware.Chain(ah.Deactivate, middleware.Logging(), middleware.Method("DELETE"))).Methods("DELETE")
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":"+addr, a.Router))
}

func main() {
	fmt.Println("Welcome to Family Tree")
	port := flag.String("port", "3333", "provide the listening port for the application. eg \"3333\"")

	app := App{}
	app.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
		"pg",
		"6379",
		"5432")

	app.Run(*port)
}
