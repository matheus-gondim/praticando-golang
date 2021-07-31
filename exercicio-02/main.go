package main

import (
	"fmt"
	"math"
)

type Points struct {
	X float64
	Y float64
}

func calculatorOfDistanceBetweenTwoPoints(p1 Points, p2 Points) float64 {
	a := math.Pow(p2.X-p1.X, 2)
	b := math.Pow(p2.Y-p1.Y, 2)
	return math.Sqrt(a + b)
}

func main() {
	var p1, p2 Points

	fmt.Print("Informe o eixo X do primeiro ponto: ")
	fmt.Scanf("%f\n", &p1.X)

	fmt.Print("Informe o eixo Y do primeiro ponto: ")
	fmt.Scanf("%f\n", &p1.Y)

	fmt.Print("Informe o eixo X do segundo ponto: ")
	fmt.Scanf("%f\n", &p2.X)

	fmt.Print("Informe o eixo Y do segundo ponto: ")
	fmt.Scanf("%f\n", &p2.Y)

	distance := calculatorOfDistanceBetweenTwoPoints(p1, p2)

	fmt.Printf("A distancia entre dois pontos é %.3f\n", distance)
}
