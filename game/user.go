package game

import (
	"bufio"
	"strings"
)

type User struct {
	Reader *bufio.Reader
}

func (u User) Move() string {
	input := readInput(u.Reader)

	return strings.TrimSpace(input)
}

func readInput(reader *bufio.Reader) string {
	text, _ := reader.ReadString('\n')
	return text
}
