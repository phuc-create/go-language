package golesson

func Closure() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
