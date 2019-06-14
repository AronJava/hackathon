package bulletscreen

import(
	"net/http"
	"text/template"
	"fmt"
)

func Index(w http.ResponseWriter, r *http.Request) {
	file, err := template.New("index.html").ParseFiles("./static/index.html")
	if err != nil {
		fmt.Println("parse file failed. err:%v", err)
		return
	}
	file.Execute(w, nil)
}
