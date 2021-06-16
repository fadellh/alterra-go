package main

import (
	"fmt"
)

func main() {

	//Challenge 01
	LuasPermukaanTabung(7, 18)
	//Challenge 02
	fmt.Println(gradeMahasiswa(45))
	////Challenge 03
	sumNumber(10)
	//Challenge 04
	startStr(5)
	//Challenge 05
	reverseStr("alterra")
}

func LuasPermukaanTabung(jari float32, tinggi float32) {
	const phi = 3.14
	var luas float32 = 2 * phi * jari * (jari + tinggi)
	fmt.Println(luas)
}

func gradeMahasiswa(grade int32) string {
	if grade > 85 {
		return "A"
	} else if grade > 75 {
		return "B"
	} else if grade > 65 {
		return "C"
	} else if grade > 55 {
		return "D"
	}
	return "Tidak Lulus"
}

func sumNumber(number int32) {
	var total int32
	for i := 1; i <= int(number); i++ {
		total = total + int32(i)
	}
	fmt.Println(total)
}

func startStr(number int32) {
	var result string = "*"
	fmt.Println(result)
	for i := 1; i < int(number); i++ {
		result = result + "*"
		fmt.Println(result)
	}
}

func reverseStr(value string) {
	var result string
	for _, r := range value {
		result = string(r) + result
	}
	fmt.Println(result)
}
