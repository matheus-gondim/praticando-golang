package main

import "fmt"

func getOlderAge(ages []int) int {
	var olderAge int

	for _, v := range ages {
		if olderAge <= v {
			olderAge = v
		}
	}

	return olderAge
}

func main() {

	ages := make([]int, 10)

	for ix := 0; ix < len(ages); ix++ {
		fmt.Printf("Informe a %v idade =>\t", ix+1)
		fmt.Scan(&ages[ix])
	}

	olderAge := getOlderAge(ages)

	fmt.Println("A maior idade informada é: ", olderAge)
}
