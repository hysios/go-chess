package game

import (
	// "math"
	"fmt"
)

type ChessType int

const (
	CTNone		ChessType = iota
    CTVehicle 
	CTHorse
	CTMinister
	CTGuardian
	CTKing
	CTConnon
	CTSolider
)

type Camp struct {
	Color int
	ID int
	POSITION int
}

type Chess struct {
	x,y int
	camp Camp
}

type Chesser interface {

}

type Movable interface {
	TryMove(cb *Chessboard, x int, y int) (bool)
	Move(cb *Chessboard, x int, y int)
}

type Vehicle struct {
	Chess
	Chesser
	Movable
}

type Horse struct {
	Chess
	Chesser
	Movable
}

type Minister struct {
	Chess
	Chesser
	Movable
}

type Guardian struct {
	Chess
	Chesser
	Movable
}

type King struct {
	Chess
	Chesser
	Movable
}

type Connon struct {
	Chess
	Chesser
	Movable
}

type Solider struct {
	Chess
	Chesser
	Movable
}

func NewChess(ct ChessType, x,y int) Chesser {
	switch ct {
	case CTVehicle:
		return &Vehicle{ 
			Chess: Chess{
				x: x, 
				y: y,
			},
		}
	case CTSolider:
		return &Solider{
			Chess: Chess{
				x: x,
				y: y,
			},
		}
	case CTConnon:
		return &Connon{
			Chess: Chess{
				x: x,
				y: y,
			},
		}	
	case CTKing:
		return &King{
			Chess: Chess{
				x: x,
				y: y,
			},
		}	
	case CTGuardian:
		return &Guardian{
			Chess: Chess{
				x: x,
				y: y,
			},
		}

	case CTMinister:
		return &Minister{
			Chess: Chess{
				x: x,
				y: y,
			},
		}

	case CTHorse:
		return &Horse{
			Chess: Chess{
				x: x,
				y: y,
			},
		}						
	default: 
		return nil
	}
}

func NewVehicle(x, y int) *Vehicle {
	return &Vehicle{ 
		Chess: Chess{
			x: x, 
			y: y,
		},
	}
}


func (v *Vehicle) TryMove(cb *Chessboard, x int, y int) bool {
	if v.x == x || v.y == y {
		return true
	} else {
		return false
	}
}

func (v *Vehicle) Move(cb *Chessboard, x int, y int) {
	if v.TryMove(cb, x, y) {
		fmt.Printf("Vehicle from %d, %d move to %d, %d", v.x, v.y, v ,y)
	} else {
		fmt.Printf("Vehicle move must in a line")
	}

}

func (m *Horse) Move(x int, y int) {
	fmt.Printf("Horse move to %d, %d", x ,y)
}

func (m *Minister) Move(x int, y int) {
	fmt.Printf("Minister move to %d, %d", x ,y)
}

func (m *Guardian) Move(x int, y int) {
	fmt.Printf("Guardian move to %d, %d", x ,y)
}

func (m *King) Move(x int, y int) {
	fmt.Printf("King move to %d, %d", x ,y)
}

func (m *Connon) Move(x int, y int) {
	fmt.Printf("Connon move to %d, %d", x ,y)
}

func (m *Solider) Move(x int, y int) {
	fmt.Printf("Solider move to %d, %d", x ,y)
}