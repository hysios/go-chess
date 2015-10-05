package web

import (
	"io"
	// "os"
	"net/http"
	"log"
	"html/template"
	"path/filepath"
	"fmt"
	"../utils"
)

var ts *template.Template
// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func handleIndex(w http.ResponseWriter, req *http.Request) {
	log.Println(TemplatePath("index.html"))
	var err error
	// log.Panic(t)

	if ts != nil {
		err = ts.ExecuteTemplate(w, "index.html", "<script>alert('you have been pwned')</script>")
	}

	if err != nil {
		log.Fatal("open template: ", err)
	}
}

var __template_dir = utils.FindDir("web/templates")

func TemplatePath(filename string) string {
	return filepath.Join(__template_dir, filename)
}

func Server() {
	var err error
	fmt.Printf("current path: %s\n", utils.FindDir("web/templates"))
	ts, err = template.ParseGlob(TemplatePath("*.html"))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	http.HandleFunc("/hello", HelloServer)
	http.HandleFunc("/", handleIndex)

	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

