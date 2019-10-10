package structs

import "math"

// Rectangle represents a shape
type Rectangle struct {
	Width  float64
	Height float64
}

// Area of rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle represents a circle
type Circle struct {
	Radius float64
}

// Area of circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Triangle is a triangle
type Triangle struct {
	Base   float64
	Height float64
}

// Area of triangle
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * .5
}

// Perimeter calculates the perimeter of the length and width provided
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}
