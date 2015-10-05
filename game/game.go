package game

import (
	"fmt"
	"github.com/nsf/termbox-go"
)
const rows = 10
const cols = 9

var DefaultGrid = [rows][cols]ChessType {
	{ CTVehicle, CTHorse, CTMinister, CTGuardian, CTKing, CTGuardian, CTMinister, CTHorse, CTVehicle},
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{ 0, CTConnon, 0, 0, 0, 0, 0, CTConnon, 0},
	{ CTSolider, 0,CTSolider, 0, CTSolider, 0, CTSolider, 0, CTSolider},
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{ CTSolider, 0,CTSolider, 0, CTSolider, 0, CTSolider, 0, CTSolider},
	{ 0, CTConnon, 0, 0, 0, 0, 0, CTConnon, 0},
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{ CTVehicle, CTHorse, CTMinister, CTGuardian, CTKing, CTGuardian, CTMinister, CTHorse, CTVehicle},
}

type Chessboard struct {
	grids [rows][cols] Chesser
}

func (cb *Chessboard) Reset() {
	for i:= 0; i < rows; i++ {
		for j:=0; j < cols; j++ {
			if ct := DefaultGrid[i][j]; ct != CTNone  {
				cb.grids[i][j] = NewChess(ct, i, j)
			} else  {
				cb.grids[i][j] = nil
			}
		}
	}
}

type Game struct {
	cb Chessboard
	players *[2]Play
}


func (g *Game) Start(players [2]*Play) {
	fmt.Printf("staring...%v", g)

	g.cb.Reset()
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetCell(1, 20, '@', 184, 240);
}

// func (g *Game) Move(chess *Chess, x, y int) (_, ok) {
// 	target := g.world.grids[x][y]
// 	originX := chess.x
// 	originY := chess.y

// 	if target && target.GetCamp() != chess.GetCamp() {
// 		g.world.grids
// 	} else if target == nil {

// 	} else {
// 		ok = error{"cant move to here"}
// 	}
// }