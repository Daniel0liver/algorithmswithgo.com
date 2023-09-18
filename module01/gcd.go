package module01

func GCD(a, b int) int {
	if b == 0 {
		return a
	}

	a, b = b, a%b

	return GCD(a, b)
}
