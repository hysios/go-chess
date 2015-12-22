package game

import "log"

// "github.com/nsf/termbox-go"

const rows = 10
const cols = 9

var DefaultGrid = [rows][cols]ChessType{
	{CTVehicle, CTHorse, CTMinister, CTGuardian, CTKing, CTGuardian, CTMinister, CTHorse, CTVehicle},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, CTConnon, 0, 0, 0, 0, 0, CTConnon, 0},
	{CTSolider, 0, CTSolider, 0, CTSolider, 0, CTSolider, 0, CTSolider},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{CTSolider, 0, CTSolider, 0, CTSolider, 0, CTSolider, 0, CTSolider},
	{0, CTConnon, 0, 0, 0, 0, 0, CTConnon, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{CTVehicle, CTHorse, CTMinister, CTGuardian, CTKing, CTGuardian, CTMinister, CTHorse, CTVehicle},
}

type Grids [rows][cols]interface{}

type Chessboard struct {
	grids Grids
}

func (cb *Chessboard) Reset() {
	origCamp := CShuai
	reveCamp := CJiang
	var camp Camp

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if ct := DefaultGrid[i][j]; ct != CTNone {
				if i < 4 {
					camp = Camp{ID: origCamp}
				} else {
					camp = Camp{ID: reveCamp}
				}

				cb.grids[i][j] = NewChess(ct, i, j, camp)
			} else {
				cb.grids[i][j] = nil
			}
		}
	}
}

type Game struct {
	cb      Chessboard
	players *[2]Play
}

func (g *Game) Start(players [2]*Play) {
	// fmt.Printf("staring...%v", g)
	log.Printf("start gaming ...%v", g)

	g.cb.Reset()
}

func (g *Game) GetGrids() Grids {
	return g.cb.grids
}
