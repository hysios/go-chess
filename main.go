package main

import (
	"fmt"
	"./game"
	"./web"
)

func main() {
	fmt.Println("chess")
	// v := game.NewVehicle(10, 15)
	v := game.NewChess(game.CTVehicle, 10, 15)
	// v := game.Vehicle{
	// 	Chess: game.Chess{
	// 		x: 10,
	// 		y: 15,
	// 	},
	// }

	fmt.Println("v := %v", v)

	players := [2]*game.Play { game.NewPlay("player1"), game.NewPlay("player2") }
	var game game.Game

	game.Start(players)

	web.Server()
	// v.Move(&cb, 10, 20)
	// var a game.Game
	// play1 := new(game.Play)
	// play2 := new(game.Play)
	// players := [2]*game.Play{play1, play2}

	// fmt.Println("%v", players)
	// a.Start(players)
}