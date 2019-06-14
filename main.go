package main

import(
	"fmt"
	"net/http"
	"os"

	"hackathon/modules"
	"hackathon/controllers"
	"hackathon/utils"
	"github.com/gorilla/sessions"
)

var routers = controllers.Routers
var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func main() {
	utils.MysqlInit()
	modules.ClientInit()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	for k, v := range routers {
		http.HandleFunc(k, v)
	}
	fmt.Println("server start.")
	http.ListenAndServe(":8089", nil)
}

func middleWare(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session_name")
		id := session.Values["userid"]
		if id == nil {
			http.Redirect(w, r, "/login", 302)
			return
		}

		r.Header.Add("User-ID", id.(string))
		h(w, r)
	}
}
