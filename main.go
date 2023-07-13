package main

import (
	"log"

	"github.com/moltenwolfcub/Forest-Game/game"
)

func main() {
	parseFlags()

	gameInstance := game.NewGame()
	if err := gameInstance.Run(); err != nil {
		log.Fatal(err)
	}
}
