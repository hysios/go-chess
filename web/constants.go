package web

import "html/template"

var funcMap = template.FuncMap{
	// The name "title" is what the function will be called in the template text.
	"calcy": calcY,
	"calcx": calcX,
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
