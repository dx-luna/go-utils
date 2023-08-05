package utilq

import (
	arrayq "github.com/dx-luna/go-utils/array"
	floatq "github.com/dx-luna/go-utils/float"
	logq "github.com/dx-luna/go-utils/log"
	stringq "github.com/dx-luna/go-utils/string"
)

// exports
var String stringq.String
var Log logq.Log
var Chalk logq.ChalkColor = logq.Chalk
var Float64 floatq.Float64
var Array arrayq.Array

func NewString(str string) stringq.String {
	return stringq.New(str)
}
func NewFloat(n float64) floatq.Float64 {
	return floatq.New(n)
}
func NewArrayString(strs []string) arrayq.Array {
	return arrayq.NewArrayString(strs)
}
