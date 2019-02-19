package utils

type IElement interface{}
type IElementSlice []IElement

func FindFirst(slice IElementSlice, cmp func(interface{}) bool) interface{} {
	for _, item := range slice {
		if cmp(item) {
			return item
		}
	}

	return nil
}

func IntRange(count int) []int {
	slice := make([]int, count)
	for i := 0; i < count; i++ {
		slice[i] = i
	}
	return slice
}
