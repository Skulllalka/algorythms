package generics

func Task2[T any](sl []T, contains func(a, b T) bool, find T) bool {
	for _, element := range sl {
		if temp := contains(element, find); temp == true {
			return true
		}
	}
	return false
}
