package utils

import "fmt"

type color int

const (
	Blank color = 30 + iota
	Red
	Green
	Yellow
	Blue
	Purple
	DarkGreen
	White
)

func Print(msg interface{}, c color) {
	fmt.Printf("\x1b[%dm%v\x1b[0m\n", c, msg)
}
