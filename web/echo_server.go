package web

import (
	"html/template"
	"io"
	"net/http"

	"github.com/hysios/go-chess/game"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

// Template provides HTML template rendering
type Template struct {
	templates *template.Template
}

// Render HTML
func (t *Template) Render(w io.Writer, name string, data interface{}) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func hello(c *echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

func index(c *echo.Context) error {
	p1 := game.NewPlay("player1")
	p2 := game.NewPlay("player2")
	players := [2]*game.Play{p1, p2}
	var game game.Game

	game.Start(players)
	return c.Render(http.StatusOK, "index.html", game.GetGrids())
}

// EchoServer is echo version of http server
func EchoServer() {
	e := echo.New()

	e.Use(mw.Logger())
	e.Use(mw.Recover())
	e.Use(mw.Gzip())
	//
	e.Static("/", "web/static")

	t := &Template{
		// Cached templates
		templates: template.Must(template.New("index.html").Funcs(funcMap).ParseFiles("web/templates/index.html")),
	}
	// t.templates.Funcs(funcMap)
	e.SetRenderer(t)

	e.Get("/", index)

	e.Run(":3000")
}
