package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mpoegel/bomberman-ai/bomberman"
	"github.com/mpoegel/bomberman-ai/strategy"
)

func main() {
	args := os.Args
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "Usage: go run ai.go <credential_file>\n")
		os.Exit(1)
	}
	player := bomberman.NewPlayer(&args[1])
	log.Printf("Loaded credentials for: %s\n", player.Username)
	strategy := strategy.NewSuicideStrategy()
	log.Printf("Using strategy: %s\n", strategy.Name)

	bomberman.PracticeGame(player, strategy)
}
