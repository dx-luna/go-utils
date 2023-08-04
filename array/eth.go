package arrayq

import "github.com/ethereum/go-ethereum/common"

type Array struct {
	ArrayString  []string
	ArrayFloat64 []float64
	ArrayUint    []float64
}

func (moreAddress Array) ToAddresses() []common.Address {
	addresses := make([]common.Address, 0)
	for _, address := range moreAddress.ArrayString {
		addresses = append(addresses, common.HexToAddress(address))
	}
	return addresses
}
