package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/hysios/go-chess/game"
	"github.com/hysios/go-chess/utils"
)

var ts *template.Template
var staticHandler http.Handler
var tpls map[string]*template.Template

var funcMap = template.FuncMap{
	// The name "title" is what the function will be called in the template text.
	"calcy": calcY,
	"calcx": calcX,
}

func HandleIndex(w http.ResponseWriter, req *http.Request) {
	log.Println(TemplatePath("index.html"))
	var err error
	var tpl *template.Template
	// log.Panic(t)

	if tpl = tpls["index.html"]; tpl == nil {
		tpl, err = template.New("index.html").Funcs(funcMap).ParseFiles(TemplatePath("index.html"))
		tpls["index.html"] = tpl
	}

	fmt.Printf("static dir: %s\n", TemplatePath("web/static"))

	players := [2]*game.Play{game.NewPlay("player1"), game.NewPlay("player2")}
	var game game.Game

	game.Start(players)

	if tpl != nil {
		err = tpl.Execute(w, game.GetGrids())
	}

	if err != nil {
		log.Fatal("open template: ", err)
	}
}

func StaticServer(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		staticHandler.ServeHTTP(w, req)
		return
	}

	HandleIndex(w, req)
}

var __template_dir = utils.FindDir("web/templates")

func TemplatePath(filename string) string {
	return filepath.Join(__template_dir, filename)
}

const (
	startTop  = 47
	startLeft = 43 + 15
	cellSize  = 77
)

func calcY(row int) int {
	return startTop + row*cellSize
}

func calcX(col int) int {
	return startLeft + col*cellSize
}

func Server() {
	var err error
	fmt.Printf("current path: %s\n", utils.FindDir("web/templates"))

	tpls = make(map[string]*template.Template)

	staticHandler = http.FileServer(http.Dir("web/static"))

	http.HandleFunc("/", StaticServer)

	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
