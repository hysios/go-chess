package game

import sr "github.com/tuvistavie/securerandom"

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
	Id      string
	Cb      Chessboard
	players *[2]Play
}

func New() (game *Game, err error) {
	if key, err := sr.UrlSafeBase64(11, false); err == nil {
		return &Game{Id: key}, nil
	}

	return nil, err
}

func (g *Game) Start(players [2]*Play) {
	// fmt.Printf("staring...%v", g)
	// log.Printf("start gaming ...%v", g)

	g.Cb.Reset()
}

// func (g *Game) UnmarshalJSON(b []byte) error {
//     var f interface{}
//     json.Unmarshal(b, &f)
//
//     m := f.(map[string]interface{})
//
//     cbmap := m["cb"]
//     v := cbmap.(map[string]interface{})
//
//     a.FooBar = v["bar"].(string)
//     a.FooBaz = v["baz"].(string)
//
//     return nil
// }
//
// func (g *Game) MarshalJSON() ([]byte, error) {
// 		json.Marshal(g.cb)
//     return []byte(`{"data":"charlie"}`), nil
// }
//
// func (g *Game) UnmarshalJSON(b []byte) error {
//     // Insert the string directly into the Data member
//     return json.Unmarshal(b, &s.Data)
// }

func (g *Game) GetGrids() Grids {
	return g.Cb.grids
}
