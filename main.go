package main

import (
	"log"

	"github.com/moltenwolfcub/Forest-Game/args"
	"github.com/moltenwolfcub/Forest-Game/game"
)

func main() {
	args.ParseFlags()

	gameInstance := game.NewGame()
	if err := gameInstance.Run(); err != nil {
		log.Fatal(err)
	}
}
