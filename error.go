package utils

import (
	"errors"
)

func CheckNoError(args ...interface{}) interface{} {
	CheckTrue(len(args) > 0 && len(args) < 3, "Unsupported argument length")

	var result interface{}
	for _, obj := range args {
		switch obj.(type) {
		case error:
			panic(obj)
		default:
			//noinspection GoNilness
			if obj != nil {
				result = obj
			}
			break
		}
	}

	return result
}

func ThrowError(message string) {
	panic(errors.New(message))
}

func CheckTrue(exp bool, message string) {
	if !exp {
		ThrowError(message)
	}
}
