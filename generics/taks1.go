package generics

func Task1[T comparable](list []T, element T) bool {
	for _, el := range list {
		if el == element {
			return true
		}
	}
	return false
}
