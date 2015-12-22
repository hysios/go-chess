package game

import (
	"fmt"
)

type ChessType int

type CampType int

const (
	CTNone ChessType = iota
	CTVehicle
	CTHorse
	CTMinister
	CTGuardian
	CTKing
	CTConnon
	CTSolider
)

const (
	CJiang CampType = iota
	CShuai
)

type Camp struct {
	Color    int
	ID       CampType
	POSITION int
}

type Chess struct {
	x, y      int
	chessType ChessType
	camp      Camp
}

type Chesser interface {
	Type() ChessType
	TypeString() string
}

type Movable interface {
	TryMove(cb *Chessboard, x int, y int) bool
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

func NewChess(ct ChessType, x, y int, camp Camp) interface{} {
	switch ct {
	case CTVehicle:
		return &Vehicle{
			Chess: Chess{
				x:    x,
				y:    y,
				camp: camp,
			},
		}
	case CTSolider:
		return &Solider{
			Chess: Chess{
				x:    x,
				y:    y,
				camp: camp,
			},
		}
	case CTConnon:
		return &Connon{
			Chess: Chess{
				x:    x,
				y:    y,
				camp: camp,
			},
		}
	case CTKing:
		return &King{
			Chess: Chess{
				x:    x,
				y:    y,
				camp: camp,
			},
		}
	case CTGuardian:
		return &Guardian{
			Chess: Chess{
				x:    x,
				y:    y,
				camp: camp,
			},
		}

	case CTMinister:
		return &Minister{
			Chess: Chess{
				x:    x,
				y:    y,
				camp: camp,
			},
		}

	case CTHorse:
		return &Horse{
			Chess: Chess{
				x:    x,
				y:    y,
				camp: camp,
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

/**
 * [func description]
 * @param  {[type]} c *Chess        [description]
 * @return {[type]}   [description]
 */
func (c *Chess) GetXY() (x, y int) {
	return c.x, c.y
}

func GetCamp(c *Chess) CampType {
	return c.camp.ID
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
		fmt.Printf("Vehicle from %d, %d move to %d, %d", v.x, v.y, v, y)
	} else {
		fmt.Printf("Vehicle move must in a line")
	}
}

func (m *Vehicle) Type() ChessType {
	return CTVehicle
}

func (m *Vehicle) TypeString() string {
	return "vehicle"
}

func (m *Horse) Move(x int, y int) {
	fmt.Printf("Horse move to %d, %d", x, y)
}

func (m *Horse) Type() ChessType {
	return CTHorse
}

func (m *Horse) TypeString() string {
	return "horse"
}

func (m *Minister) Move(x int, y int) {
	fmt.Printf("Minister move to %d, %d", x, y)
}

func (m *Minister) Type() ChessType {
	return CTMinister
}

func (m *Minister) TypeString() string {
	return "minister"
}

func (m *Guardian) Move(x int, y int) {
	fmt.Printf("Guardian move to %d, %d", x, y)
}

func (m *Guardian) Type() ChessType {
	return CTGuardian
}

func (m *Guardian) TypeString() string {
	return "guardian"
}

func (m *King) Move(x int, y int) {
	fmt.Printf("King move to %d, %d", x, y)
}

func (m *King) Type() ChessType {
	return CTKing
}

func (m *King) TypeString() string {
	return "king"
}

func (m *Connon) Move(x int, y int) {
	fmt.Printf("Connon move to %d, %d", x, y)
}

func (m *Connon) Type() ChessType {
	return CTConnon
}

func (m *Connon) TypeString() string {
	return "connon"
}

func (m *Solider) Move(x int, y int) {
	fmt.Printf("Solider move to %d, %d", x, y)
}

func (m *Solider) Type() ChessType {
	return CTSolider
}

func (m *Solider) TypeString() string {
	return "solider"
}

func (c *Chess) Camp() int {
	return int(c.camp.ID)
}
