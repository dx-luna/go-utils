package logq

import "fmt"

type ChalkColor struct {
	Reset  string
	Red    string
	Green  string
	Yellow string
	Blue   string
	Purple string
	Cyan   string
	Gray   string
	White  string
}

var Chalk = ChalkColor{
	Reset:  "\033[0m",
	Red:    "\033[31m",
	Green:  "\033[32m",
	Yellow: "\033[33m",
	Blue:   "\033[34m",
	Purple: "\033[35m",
	Cyan:   "\033[36m",
	Gray:   "\033[37m",
	White:  "\033[97m",
}
var chalk = map[string]string{
	"reset":  "\033[0m",
	"red":    "\033[31m",
	"green":  "\033[32m",
	"yellow": "\033[33m",
	"blue":   "\033[34m",
	"purple": "\033[35m",
	"cyan":   "\033[36m",
	"cray":   "\033[37m",
	"white":  "\033[97m",
}

func (blank Log) PrintColor(color string, input any) {
	// type PrintL struct {
	// 	Color string
	// 	Input any
	// 	Reset string
	// }
	// genr := PrintL{
	// 	Color: color,
	// 	Input: input,
	// 	Reset: Chalk.Reset,
	// }
	// fmt.Println(genr)
	fmt.Println(color, input, Chalk.Reset)
}
