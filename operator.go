package utils

func Ternary(condition bool, valueOk interface{}, valueOtherwise interface{}) interface{} {
	if condition {
		return valueOk
	}
	return valueOtherwise
}
