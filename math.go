package utils

func selectIntVal(cmp func(int, int) bool, a int, b ...int) int {
	for _, v := range b {
		if cmp(a, v) {
			a = v
		}
	}

	return a
}

func MaxInt(a int, b ...int) int {
	return selectIntVal(func(a, b int) bool { return a > b }, a, b...)
}

func MinInt(a int, b ...int) int {
	return selectIntVal(func(a, b int) bool { return a < b }, a, b...)
}

func AbsInt(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func PowInt(a, b int) int {
	result := 1
	for i := 0; i < b; i++ {
		result = result * a
	}
	return result
}

func Sign(val int) int {
	if val == 0 {
		return 0
	}
	return val / AbsInt(val)
}
