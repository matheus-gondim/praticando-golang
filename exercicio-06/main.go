package main

import (
	"fmt"
	"math"
)

const WeightByCubicCm float64 = 7.8

type washers struct {
	Height       float64
	ExtDiam      float64
	IntDiam      float64
	Qt           int
	FreightPerKg float64
}

func shippingPrice(w washers, weight float64) string {
	extRay, intRay := w.ExtDiam/2, w.IntDiam/2

	vol := math.Pi * w.Height * (math.Pow(extRay, 2) - math.Pow(intRay, 2))

	totalWeight := vol * float64(w.Qt) * weight

	sp := totalWeight * w.FreightPerKg / 1000

	return fmt.Sprintf("R$%.2f", sp)
}

func main() {
	var inputData washers

	fmt.Printf("Digite o peso da arruela: ")
	fmt.Scanf("%f\n", &inputData.Height)
	fmt.Printf("Digite o diâmetro externo da arruela: ")
	fmt.Scanf("%f\n", &inputData.ExtDiam)
	fmt.Printf("Digite o diâmetro interno da arruela: ")
	fmt.Scanf("%f\n", &inputData.IntDiam)
	fmt.Printf("Digite a quantidade de arruelas: ")
	fmt.Scanf("%d\n", &inputData.Qt)
	fmt.Printf("Digite o frete por quilograma: ")
	fmt.Scanf("%f\n", &inputData.FreightPerKg)

	fmt.Printf("Preço total do frete: %s\n", shippingPrice(inputData, WeightByCubicCm))
}
