package main

import (
	"fmt"

	utilq "github.com/dx-luna/go-utils"
)

var jsonf = map[string]string{
	"name":     "john",
	"lastname": "doe",
}

func main() {
	utilq.Log.ClearLog()
	initial := utilq.String.New("554.58")
	fixed := initial.ToFixed(18)
	fmt.Println(initial, fixed, initial.ToInterface())
	init := utilq.String.New("22.56").ToFixed(16)
	fmt.Println(init)
	utilq.Log.PrintColor(utilq.Chalk.Green, "Win: 441")
	fmt.Println("Lost :", utilq.Chalk.Red, "-455", utilq.Chalk.Reset)
	utilq.Log.PrintJson(jsonf)
	fmt.Println(utilq.Float64.New(0.12).ToEther(), float64(2/3))
	fmt.Println(utilq.Float64.New(5).Add(6).Sub(1).Div(10))
	utilq.Log.PrintColor(utilq.Chalk.Yellow, "hello Yellow ol")

	var newStr []string = make([]string, 2)
	newStr[0] = "45.666"
	newStr[1] = "12.77"
	fmt.Println(utilq.Array.NewArrayString(newStr).ToFloat64())
}
