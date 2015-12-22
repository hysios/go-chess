package main

import (
	"github.com/hysios/go-chess/web"
)

/**
 * [main description]
 * @return {[type]} [description]
 */
func main() {
	web.EchoServer()

	foo()
}

func foo() (b int) {
	b = 1
	return b
}
