package log

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

type Log struct {
	val any
}

func (blank Log) ClearLog() {
	fmt.Printf("\x1bc")
}

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
func (blank Log) ClearTerminal() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}
func (blank Log) ViewJson(result ...interface{}) string {
	res, _ := json.MarshalIndent(result, "", "\t")
	return string(res)
}
func (blank Log) PrintJson(result ...interface{}) {
	res, _ := json.MarshalIndent(result, "", "\t")
	fmt.Println(string(res))
}
