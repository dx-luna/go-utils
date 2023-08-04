package floatq

import (
	"fmt"
	"math/big"

	"github.com/chenzhijie/go-web3/utils"
)

var Utils *utils.Utils = utils.NewUtils()

type Float64 struct {
	val float64
}

func New(num float64) Float64 {
	return Float64{val: num}
}
func (num Float64) Val() float64 {
	return num.val
}
func (blank Float64) New(num float64) Float64 {
	return Float64{val: num}
}

func (num Float64) ToString() string {
	return fmt.Sprintf("%.18f", num)
}
func ToFixedFloat64(a float64, b string) string {
	return fmt.Sprintf("%."+b+"f", a)
}
func (blank Float64) ToFixed(a float64, b string) string {
	return fmt.Sprintf("%."+b+"f", a)
}
func (num Float64) ToWei() *big.Int {
	str := ToFixedFloat64(num.val/1e18, "18")
	return Utils.ToWei(str)
}
func (num Float64) ToEther() *big.Int {
	str := ToFixedFloat64(num.val, "18")
	return Utils.ToWei(str)
}
func (d Float64) SubFee(e float64) float64 {
	return d.val - (e * d.val)
}
