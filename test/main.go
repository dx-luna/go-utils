package main

import (
	"fmt"

	utilq "github.com/dx-luna/go-utils"
	arrayq "github.com/dx-luna/go-utils/array"
	cliq "github.com/dx-luna/go-utils/cli"
	logq "github.com/dx-luna/go-utils/log"
)

var jsonf = map[string]string{
	"fname":    "john",
	"lastname": "doe",
}

func main() {
	logq.Clear()
	initial := utilq.String.New("554.58")
	fixed := initial.ToFixed(18)
	fmt.Println(initial, fixed, initial.ToInterface())
	init := utilq.String.New("22.56").ToFixed(16)
	fmt.Println(init)
	logq.PrintColor(utilq.Chalk.Green, "Win: 441")
	fmt.Println("Lost :", utilq.Chalk.Red, "-455", utilq.Chalk.Reset)
	logq.PrintJson(jsonf)
	fmt.Println(utilq.Float64.New(0.12).ToEther(), float64(2/3))
	fmt.Println(utilq.Float64.New(5).Add(6).Sub(1).Div(10))
	logq.PrintColor("green", "hello Yellow ol")

	logq.Clear()

	var newStr []string = make([]string, 2)
	newStr[0] = "45.666"
	newStr[1] = "12.776"
	fmt.Println(arrayq.NewArrayString(newStr).ToFloat64())
	logq.Warn("Yellow", "green a")

	// cliq.WriteFile("../cli/test.json", jsonf)
	fmt.Println(cliq.ReadDir("../cli"))
}
