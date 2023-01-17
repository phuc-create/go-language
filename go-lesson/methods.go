package golesson

type Square struct {
	X, Y float64
}
type Triangle struct {
	B, C float64
}

func (sq Square) Area() float64 {
	return sq.X * sq.Y
}

func (tg Triangle) Area() float64 {
	return (tg.B + tg.C) * 2
}
