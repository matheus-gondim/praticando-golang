package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
)

func binaryToDecimal(binary string) (decimal float64, err error) {
	message := fmt.Sprintf("Número inválido: %s", binary)

	if _, err := strconv.Atoi(binary); err != nil {
		return 0, errors.New(message)
	}

	for ix, value := range binary {
		strValue := string(value)
		number, _ := strconv.Atoi(strValue)
		exponent := len(binary) - ix - 1

		if number == 1 {
			decimal += math.Pow(2, float64(exponent))
		} else if number > 1 {
			return 0, errors.New(message)
		}
	}

	return
}

func main() {
	decinal, err := binaryToDecimal("111100010101011101")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(decinal)
}
