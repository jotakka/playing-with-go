package exercism

import "math"

type Number struct {
	a float64 // real
	b float64 // complex
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
		a: n1.a*n2.a - n1.b*n2.b,
		b: n1.b*n2.a + n1.a*n2.b,
	}
}

func (n Number) Times(factor float64) Number {
	return Number{
		a: factor * n.a,
		b: factor * n.b,
	}
}

func (n1 Number) Divide(n2 Number) Number {
	return Number{
		a: (n1.a*n2.a + n1.b*n2.b) / (n2.a*n2.a + n2.b*n2.b),
		b: (n1.b*n2.a - n1.a*n2.b) / (n2.a*n2.a + n2.b*n2.b),
	}
}

func (n Number) Conjugate() Number {
	return Number{
		a: n.a,
		b: -n.b,
	}
}

func (n Number) Abs() float64 {
	return math.Sqrt(n.a*n.a + n.b*n.b)
}

func (n Number) Exp() Number {
	abs := math.Exp(n.a)
	return Number{
		a: abs * math.Cos(n.b),
		b: abs * math.Sin(n.b),
	}
}
