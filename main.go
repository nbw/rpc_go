package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/nbw/rps/game"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		log.Fatal("Not enough arguements")
	}

	rounds := args[0]

	r, err := strconv.Atoi(rounds)
	if err != nil {
		log.Fatal(fmt.Sprintf("%q is not a number.\n", rounds))
	}

	game.Start(r)
}
