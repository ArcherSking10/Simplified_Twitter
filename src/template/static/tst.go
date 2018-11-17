package main

import (
	"html/template"
	"log"
	"net/http"
	// "fmt"
	"github.com/gorilla/mux"
)

var tmpls = template.Must(template.ParseFiles("templates/index.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	if err := tmpls.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// type MyMux struct {
// }

func main() {
	// mux := &MyMux{}
	r := mux.NewRouter()
	r.HandleFunc("/", Index)

	r.PathPrefix("/styles/").Handler(http.StripPrefix("/styles/",
		http.FileServer(http.Dir("templates/styles/"))))

	http.Handle("/", r)
	log.Fatalln(http.ListenAndServe(":9000", nil))
}
