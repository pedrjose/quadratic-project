
package quadratic

import (
	"fmt"
	"math"
)

func Discriminant(a float64, b float64, c float64) float64 {
	discriminant := (b * b) - (4 * a * c)

	return discriminant
}

func Bhaskara(delta float64, a float64, b float64) string {
	if delta < 0 {
		fmt.Println("No real solutions!")
		return "Finished Algorithm!"
	}

	fmt.Println("2 real solutions:")

	x1 := (-b + math.Sqrt(delta)) / (2 * a)
	x2 := (-b - math.Sqrt(delta)) / (2 * a)

	fmt.Println("x¹ =", x1, "and x² =", x2)

	return "Finished Algorithm!"
}

