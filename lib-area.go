package goarea

import "math"

// Pi is a numerical ratio defined by the ratio between the perimeter of a circle and its diameter
const Pi = 3.1416

// Circle is responsible for calculating the area of the circuference
func Circle(radius float64) float64 {
	return math.Pow(radius, 2) * Pi
}

// Rectangle is responsible for calculating the area of rectangle
func Rectangle(base, height float64) float64 {
	return base * height
}

// Its is not visible
func _Triangle(base, height float64) float64 {
	return (base * height) / 2
}
