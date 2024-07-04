package math

type Math struct {
	A int
	B int
}

func (m Math) Add() int {
	return m.A + m.B
}

type MathPrivateFields struct {
	a int
	b int
}

func NewMathPrivateFields(a, b int) MathPrivateFields {
	return MathPrivateFields{a: a, b: b}
}

func (m MathPrivateFields) Add() int {
	return m.a + m.b
}

type mathPrivate struct {
	a int
	b int
}

func NewMathPrivate(a, b int) mathPrivate {
	return mathPrivate{a: a, b: b}
}

func (m mathPrivate) Add() int {
	return m.a + m.b
}
