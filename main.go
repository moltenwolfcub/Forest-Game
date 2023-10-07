package main

import (
	"log/slog"

	"github.com/moltenwolfcub/Forest-Game/args"
	"github.com/moltenwolfcub/Forest-Game/game"
)

func main() {
	args.ParseFlags()
	if args.ProfilerFlag {
		stop := args.Profile()
		defer stop()
	}

	gameInstance, err := game.NewGame()
	if err != nil {
		panic(err)
	}

	if err := gameInstance.Run(); err != nil {
		slog.Error("Failed to execute tick: " + err.Error())
		return
	}
}
