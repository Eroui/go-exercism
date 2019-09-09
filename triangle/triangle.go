// Package triangle is a utility package for triangles geometric shapes
package triangle

import "math"

// Kind is a type to define the kind of a triangle
type Kind string

// Constants defining Kinds of triangles
const (
	NaT Kind = "Not A Triangle"
	Equ Kind = "Equilateral Triangle"
	Iso Kind = "Isosceles Triangle"
	Sca Kind = "Scalene Triangle"
)

// KindFromSides Determines the kind of a triangle
func KindFromSides(a, b, c float64) Kind {
	if !isSide(a) || !isSide(b) || !isSide(c) {
		return NaT
	}
	if !isTriangle(a, b, c) {
		return NaT
	} else if a == b && a == c && b == c {
		return Equ
	} else if a == b || a == c || b == c {
		return Iso
	} else if a != b && a != c && b != c {
		return Sca
	}
	return NaT
}

func isTriangle(a, b, c float64) bool {
	return a+b >= c && a+c >= b && b+c >= a
}

func isSide(a float64) bool {
	return a > 0 && !math.IsInf(a, 1) && !math.IsNaN(a)
}
