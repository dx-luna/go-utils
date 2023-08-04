package floatq

func (a Float64) Add(b float64) Float64 {
	return Float64{val: a.val + b}
}
func (a Float64) Sub(b float64) Float64 {
	return Float64{val: a.val - b}
}
func (a Float64) Mul(b float64) Float64 {
	return Float64{val: a.val * b}
}
func (a Float64) Div(b float64) Float64 {
	return Float64{val: a.val / b}
}

func (a Float64) MAdd(b Float64) Float64 {
	return Float64{val: a.val + b.val}
}
func (a Float64) MSub(b Float64) Float64 {
	return Float64{val: a.val - b.val}
}
func (a Float64) MMul(b Float64) Float64 {
	return Float64{val: a.val * b.val}
}
func (a Float64) MDiv(b Float64) Float64 {
	return Float64{val: a.val / b.val}
}
