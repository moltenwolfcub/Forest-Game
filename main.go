package main

import (
	"log/slog"
	"os"

	"github.com/moltenwolfcub/Forest-Game/args"
	"github.com/moltenwolfcub/Forest-Game/game"
)

func main() {
	args.ParseFlags()

	gameInstance, err := game.NewGame()
	if err != nil {
		panic(err)
	}

	if err := gameInstance.Run(); err != nil {
		slog.Error("Failed to execute tick: " + err.Error())
		os.Exit(1)
	}
}
