package floatq

func (s Float64) DecimalOf(req int) float64 {
	listDecimal := map[int]float64{
		1:  10,
		2:  100,
		3:  1_000,
		4:  10_000,
		5:  100_000,
		6:  1_000_000,
		7:  10_000_000,
		8:  100_000_000,
		9:  1_000_000_000,
		10: 10_000_000_000,
		11: 100_000_000_000,
		12: 1_000_000_000_000,
		13: 10_000_000_000_000,
		14: 100_000_000_000_000,
		15: 1_000_000_000_000_000,
		16: 10_000_000_000_000_000,
		17: 100_000_000_000_000_000,
		18: 1_000_000_000_000_000_000,
	}
	return listDecimal[req]
}
