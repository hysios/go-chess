package web

import (
	"html/template"

	"github.com/go-martini/martini"
	"github.com/hysios/go-chess/game"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/staticbin"
)

// Asset is bindata Asset Func type wrapper
type asset func(string) ([]byte, error)

// FuncMapArray is FuncMap Array
type FuncMapArray []template.FuncMap

var funcs = FuncMapArray{funcMap}

// Server2 Launcher
func Server2(a asset) {
	m := martini.Classic()
	m.Use(staticbin.Static("web/static", a))
	m.Use(render.Renderer(render.Options{
		Directory:  "web/templates",
		Extensions: []string{".tmpl", ".html"},
		Funcs:      funcs, // Specify helper function maps for templates to access.
	}))

	m.Get("/", func(r render.Render) {
		p1 := game.NewPlay("player1")
		p2 := game.NewPlay("player2")
		players := [2]*game.Play{p1, p2}
		var game game.Game

		game.Start(players)

		r.HTML(200, "index", game.GetGrids())
	})
	m.Run()
}
