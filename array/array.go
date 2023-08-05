package arrayq

import (
	"log"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
)

type Array struct {
	ArrayString  []string
	ArrayFloat64 []float64
	ArrayUint    []uint
}

func NewArrayString(strs []string) Array {
	return Array{ArrayString: strs}
}
func NewArrayFloat64(numbs []float64) Array {
	return Array{ArrayFloat64: numbs}
}
func NewArrayUInt(numbs []uint) Array {
	return Array{ArrayUint: numbs}
}
func StringToFloat64(s string) float64 {
	res, err := strconv.ParseFloat(s, 0)
	if err != nil {
		log.Fatalln("error convert from string to float", err)
	}
	return res
}
func SliceArray(arrs [][]string, count int) [][][]string {
	res := make([][][]string, 0)
	_len := len(arrs)
	for i := 0; i < _len; i += count {
		j := i + count
		if j > _len {
			j = _len
		}
		res = append(res, arrs[i:j])
	}
	return res
}
func (moreAddress Array) ToAddresses() []common.Address {
	addresses := make([]common.Address, 0)
	for _, address := range moreAddress.ArrayString {
		addresses = append(addresses, common.HexToAddress(address))
	}
	return addresses
}
func (moreString Array) ToFloat64() []float64 {
	results := make([]float64, len(moreString.ArrayString))
	for key, value := range moreString.ArrayString {
		results[key] = StringToFloat64(value)
	}
	return results
}
