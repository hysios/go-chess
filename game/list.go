package game

import (
	"bufio"
	"container/list"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
)

var gameList *list.List

const gamesFile = "games.xml"

func Find(id string) (game *Game, err error) {
	for e := gameList.Front(); e != nil; e = e.Next() {
		if game, ok := e.Value.(*Game); ok {
			if game.Id == id {
				return game, nil
			}
		}
	}

	return nil, fmt.Errorf("game: can't found this %s id's game Object", id)
}

func init() {
	gameList = list.New()
	loadData()
}

func loadData() {
	f, err := os.OpenFile(gamesFile, os.O_RDONLY, 0777)
	defer f.Close()
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}

	dec := xml.NewDecoder(bufio.NewReader(f))

	for {
		var g Game
		if err := dec.Decode(&g); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		gameList.PushBack(&g)
	}
}

func Save() {
	f, err := os.Create(gamesFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	// }

	w := bufio.NewWriter(f)
	enc := xml.NewEncoder(w)

	for e := gameList.Front(); e != nil; e = e.Next() {
		if v := e.Value.(*Game); v != nil {
			fmt.Println("%v", v)
			if err := enc.Encode(v); err != nil {
				log.Fatal(err)
			}
		}
	}
	w.Flush()
}

func Push(g *Game) {
	gameList.PushBack(g)
}
