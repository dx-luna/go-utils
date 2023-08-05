package stringq

import (
	"strconv"

	"github.com/ethereum/go-ethereum/common"
)

type String struct {
	value string
}

func New(str string) String {
	return String{value: str}
}
func (blank String) New(str string) String {
	return String{value: str}
}
func (str String) ToFixed(fixed int) float64 {
	res, _ := strconv.ParseFloat(str.value, fixed)
	return res
}
func (str String) ToFloat64(fixed int) float64 {
	res, _ := strconv.ParseFloat(str.value, fixed)
	return res
}
func (str String) ToInt() int {
	return int(str.ToFloat64(0))
}
func (str String) ToUint() uint {
	return uint(str.ToFloat64(0))
}
func (str String) ToInterface() interface{} {
	return str.value
}
func (str String) ToBytes() []byte {
	return []byte(str.value)
}
func (str String) ToAddress() common.Address {
	return common.HexToAddress(str.value)
}
