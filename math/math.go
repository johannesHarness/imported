package math

func Add(values ...int) int {
	res := 0
	for v := range values {
		res += v
	}
	return res
}
