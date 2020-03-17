package teststruct

import "math"

type Triangle struct {
	A int
	B int
	C int
}

func (t *Triangle) Area() int {
	s := t.Girth()
	return int(math.Sqrt(float64(s * (s - t.A) * (s - t.B) * (s - t.C))))
}

func (t *Triangle) Girth() int {
	return t.A + t.B + t.C
}
