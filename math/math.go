package math

func Sum(values ...int) int {
	res := 0
	for _, v := range values {
		res += v
	}
	return res
}
