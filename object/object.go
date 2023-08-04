package objectq

type Object struct {
	value any
}

func ObjectKeys(obj map[string]string) []string {
	keys := make([]string, len(obj))
	i := 0
	for K := range obj {
		keys[i] = K
		i++
	}
	return keys
}
