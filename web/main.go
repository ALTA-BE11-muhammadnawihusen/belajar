package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Ini adalah isi dari sebuah website")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Halo")
	})

	http.HandleFunc("/index", Index)

	http.HandleFunc("/template", func(w http.ResponseWriter, r *http.Request) {
		mapp := map[string]string{
			"Name":  "Muhammad Nawi Husen",
			"Kerja": "Backend Engineer",
		}

		t, err := template.ParseFiles("./web/temp.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(w, mapp)
	})

	fmt.Println("starting web server at http://localhost:8000/")
	http.ListenAndServe(":8000", nil)
}
