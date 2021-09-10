package main

import (
	"fmt"
	"math"
)

type Dto struct {
	diametroTubo    float64
	compArranhador  float64
	grossCordaSisal float64
}

func calculaMetragem(dto Dto) float64 {
	return (dto.compArranhador / dto.grossCordaSisal) * (2 * math.Pi * (dto.diametroTubo / 2))
}

func main() {
	var medidas Dto
	fmt.Printf("Diâmetro do Tubo Kraft: ")
	fmt.Scanf("%f\n", &medidas.diametroTubo)
	fmt.Printf("Comprimento do Arranhador: ")
	fmt.Scanf("%f\n", &medidas.compArranhador)
	fmt.Printf("Grossura da Corda de Sisal: ")
	fmt.Scanf("%f\n", &medidas.grossCordaSisal)

	resultado := calculaMetragem(medidas)
	fmt.Printf("Será necessário %.1f Metros de corda de sisal por arranhador\n", resultado)
}
