package main

import(
	"fmt"
	"net/http"

	"hackathon/modules"
	"hackathon/controllers"
)

var routers = controllers.Routers

func main() {
	modules.ClientInit()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	for k, v := range routers {
		http.HandleFunc(k, v)
	}
	fmt.Println("server start.")
	http.ListenAndServe(":8089", nil)
}
