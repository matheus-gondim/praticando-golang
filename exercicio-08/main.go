package main

import "fmt"

func SplitString(str string) []string {
	stringLen := len(str)

	if module := stringLen % 2; module != 0 {
		str += "_"
	}

	res := []string{}
	for i := 0; i < stringLen; i += 2 {
		res = append(res, str[i:i+2])
	}

	return res
}

func main() {
	fmt.Println(SplitString("abc"))
}
