package objectq

type BigMap struct {
	AllKeys  []int
	LastKey  int
	MAX_SIZE int
	Data     map[int]map[int]any
}

func NewBigMap() BigMap {

	return BigMap{
		LastKey:  1,
		AllKeys:  []int{1},
		MAX_SIZE: 1e6 * 9, // 9 million
		Data:     make(map[int]map[int]any, 1e6*9),
	}
}
func (s BigMap) Set(req int, value any) error {
	if len(s.Data[s.LastKey]) < s.MAX_SIZE {
		s.Data[s.LastKey][req] = value
		return nil
	}
	newMapKey := s.NextKey()
	s.Data[newMapKey] = make(map[int]any, 1e6*9)
	s.Data[newMapKey][0] = value
	s.LastKey = newMapKey
	return nil
}
func (s BigMap) Get(req int) any {
	var res map[int]any
	for key := range s.AllKeys {
		if s.Data[key][req] != nil {
			return s.Data[key][req]
		}
	}
	return res
}
func (s BigMap) NextKey() int {
	return s.LastKey + 1
}

func (s BigMap) Has(req int) bool {
	for key := range s.AllKeys {
		if s.Data[key][req] != nil {
			return true
		}
	}
	return false
}
