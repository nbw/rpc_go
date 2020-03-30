package game

import (
	"bufio"
	"fmt"
	"os"
)

const WON = "won"
const TIED = "tied"
const LOST = "lost"

func Start(rounds int) {
	results := map[string]int{
		WON:  0,
		TIED: 0,
		LOST: 0,
	}

	fmt.Printf("%v\n", rounds)
	fmt.Println("Rock.. Paper.. Scissors..")

	reader := bufio.NewReader(os.Stdin)

	u := User{reader}
	e := Enemy{}

	var uMove, eMove string
	for i := 0; i < rounds; i++ {
		fmt.Print("Your move: ")
		uMove = u.Move()

		eMove = e.Move()
		fmt.Print("Enemy move: ")
		fmt.Println(eMove)

		fmt.Printf("You %s.\n\n", calcResults(&results, uMove, eMove))
	}

	fmt.Println(results)
}

func calcResults(r *map[string]int, p1, p2 string) string {
	switch cm := counterMove(p1); {
	case cm == p2:
		(*r)[WON]++
		return WON
	case p1 == p2:
		(*r)[TIED]++
		return TIED
	default:
		(*r)[LOST]++
		return LOST
	}
}

func counterMove(move string) string {
	switch move {
	case "rock":
		return "scissors"
	case "paper":
		return "rock"
	case "scissors":
		return "paper"
	default:
		return ""
	}
}
