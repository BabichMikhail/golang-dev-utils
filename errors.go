package utils

import (
	"errors"
)

func AssertNoError(args ...interface{}) interface{} {
	AssertTrue(len(args) > 0 && len(args) < 3, "Unsupported argument length")

	for _, obj := range args {
		switch obj.(type) {
		case error:
			err := args[0]
			if err != nil {
				panic(err)
			}
			break
		default:
			break
		}
	}

	var result interface{}
	if len(args) == 2 {
		result = args[0]
	}

	return result
}

func ThrowError(message string) {
	panic(errors.New(message))
}

func AssertTrue(exp bool, message string) {
	if !exp {
		ThrowError(message)
	}
}
