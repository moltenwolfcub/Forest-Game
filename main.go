package main

import (
	"log"

	"github.com/moltenwolfcub/Forest-Game/game"
)

func main() {
	gameInstance := game.Game{}
	if err := gameInstance.Run(); err != nil {
		log.Fatal(err)
	}
}
