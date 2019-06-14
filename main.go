package main

import(
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/sessions"

	"hackathon/modules/cl"
	"hackathon/controllers"
	"hackathon/utils"
	"hackathon/modules/session"
)

var routers = controllers.Routers

func main() {
	utils.MysqlInit()
	cl.ClientInit()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	for k, v := range routers {
		http.HandleFunc(k, v)
	}
	fmt.Println("server start.")
	http.ListenAndServe(":8089", nil)
}

func middleWare(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s, _ := session.Store.Get(r, "session_name")
		id := s.Values["userid"]
		if r.RequestURL != "/account/login" && id == nil {
			http.Redirect(w, r, "/account/login", 302)
			return
		}

		if id != nil {
			r.Header.Add("User-ID", id.(string))
		}

		h(w, r)
	}
}
