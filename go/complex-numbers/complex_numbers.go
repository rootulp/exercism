package complexnumbers

// Define the Number type here.
type Number struct {
	a float64
	b float64
}

func (n Number) Real() float64 {
	return n.a
}

func (n Number) Imaginary() float64 {
	return n.b
}

func (n1 Number) Add(n2 Number) Number {
	return Number{
		a: n1.a + n2.a,
		b: n1.b + n2.b,
	}
}

func (n1 Number) Subtract(n2 Number) Number {
	return Number{
		a: n1.a - n2.a,
		b: n1.b - n2.b,
	}
}

func (n1 Number) Multiply(n2 Number) Number {
	return Number{
		a: (n1.a * n2.a) - (n1.b * n2.b),
		b: (n1.b * n2.a) + (n1.a * n2.b),
	}
}

func (n Number) Times(factor float64) Number {
	panic("Please implement the Times method")
}

func (n1 Number) Divide(n2 Number) Number {
	panic("Please implement the Divide method")
}

func (n Number) Conjugate() Number {
	panic("Please implement the Conjugate method")
}

func (n Number) Abs() float64 {
	panic("Please implement the Abs method")
}

func (n Number) Exp() Number {
	panic("Please implement the Exp method")
}
