package main

import "fmt"

func milesToKm(m float32) float32 {
	return m * 1.609
}

func main() {
	var miles float32
	fmt.Scanf("%f\n", &miles)
	km := milesToKm(miles)
	fmt.Println(km)
}
