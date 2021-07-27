package main

import "fmt"

type Ponto struct {
	X float32
	Y float32
}

func differenceBetweenTwoPoints() {}

func main() {
	var p1 Ponto
	var p2 Ponto

	fmt.Print("Informe o eixo X do primeiro ponto: ")
	fmt.Scanf("%f\n", &p1.X)

	fmt.Print("Informe o eixo Y do primeiro ponto: ")
	fmt.Scanf("%f\n", &p1.Y)

	fmt.Print("Informe o eixo X do segundo ponto: ")
	fmt.Scanf("%f\n", &p2.X)

	fmt.Print("Informe o eixo Y do segundo ponto: ")
	fmt.Scanf("%f\n", &p2.Y)

	fmt.Println(p1)
}
