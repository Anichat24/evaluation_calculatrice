package calculatrice

import (
	"fmt"
)

func Calculatrice(calcul []rune) int {
	var num1, num2 float64
	var op string
	//demander un chiffre

	switch op {
	case "+":
		fmt.Printf("%f %s %f = %f", num1, op, num2, num1+num2)
	case "-":
		fmt.Printf("%f %s %f = %f", num1, op, num2, num1-num2)
	case "*":
		fmt.Printf("%f %s %f = %f", num1, op, num2, num1*num2)

	default:
		fmt.Printf("nope")

	}
	return Calculatrice(calcul)
}
