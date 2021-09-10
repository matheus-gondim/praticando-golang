package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

}

func calculateMedia(data []float64) (float64, error) {
	var res float64
	return res, nil
}

func parseData(input []string) ([]float64, error) {
	var pd []float64
	for _, value := range input {
		value = strings.ReplaceAll(value, ",", ".")
		valueInFloat64, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return pd, fmt.Errorf("Incorrect Data: failed processing value %v", v)
		}
		pd = append(pd, valueInFloat64)
	}
	return pd, nil
}
