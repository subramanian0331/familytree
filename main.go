package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
	rg "github.com/redislabs/redisgraph-go"
	"github.com/subramanian0331/familytree/handlers"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize(user, password, dbname string, host string) {
	a.Router = mux.NewRouter()
	h := handlers.NewBaseHandler(host)
	a.initializeRoutes(h)
}

func (a *App) initializeRoutes(h *handlers.BaseHandler) {
	a.Router.HandleFunc("/health/", h.HealthCheck).Methods("GET")
	a.Router.HandleFunc("/addMember/", h.AddMember).Methods("POST")
	a.Router.HandleFunc("/getMember/{memberId}", h.GetMember).Methods("GET")
	// a.Router.HandleFunc("/GetMember", h.GetMember).Methods("POST")
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":"+addr, a.Router))
}

func main() {
	fmt.Println("Welcome to Family Tree")
	port := flag.String("listening port", "3333", "provide the listening port for the application. eg \"3333\"")
	flag.Parse()

	key := "Secret-session-key" // Replace with your SESSION_SECRET or similar
	maxAge := 86400 * 30        // 30 days
	isProd := false             // Set to true when serving over https

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = isProd

	gothic.Store = store

	goth.UseProviders(
		google.New("49422998462-fomnp86pputufo8tpaqcr1rdhq43f8pe.apps.googleusercontent.com", "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", "http://localhost:3000/auth/google/callback", "email", "profile"),
	)

	app := App{}
	app.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
		"127.0.0.1:6379")

	app.Run(*port)

	conn, _ := redis.Dial("tcp", "127.0.0.1:6379")
	g1 := rg.GraphNew("test", conn)
	john := rg.Node{
		Label: "person",
		Properties: map[string]interface{}{
			"name":   "John Doe",
			"age":    33,
			"gender": "male",
			"status": "single",
		},
	}
	g1.AddNode(&john)
	japan := rg.Node{
		Label: "country",
		Properties: map[string]interface{}{
			"name": "Japan",
		},
	}
	g1.AddNode(&japan)

	edge := rg.Edge{
		Source:      &john,
		Relation:    "visited",
		Destination: &japan,
	}
	g1.AddEdge(&edge)
	g1.Commit()
	conn.Close()
	fmt.Println("i am here")
	conn2, _ := redis.Dial("tcp", "127.0.0.1:6379")
	g2 := rg.GraphNew("test", conn2)
	query := `MATCH (p:person)-[v:visited]->(c:country)
           RETURN p.name, p.age, c.name`
	// query2 := fmt.Sprintf(`MATCH (x) WHERE x.name = %s RETURN x`, "John Doe")
	query2 := "MATCH p = (:person)-[:visited]->(:country) RETURN p"

	g2.Query(query2)
	result, err1 := g2.Query(query)
	fmt.Println("err:", err1)
	result.PrettyPrint()
	for result.Next() {
		val := result.Record()
		// p, ok := val.Get("name")
		p, ok := val.GetByIndex(0).(string)
		fmt.Printf("%v %v\n", p, ok)
	}

}
