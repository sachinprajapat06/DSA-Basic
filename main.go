package main

import (
	star "dsa/starPatterns"
	"fmt"
)

func main() {
	//starting
	fmt.Println("Started")

	// // *************************************numeric operations************************************

	// // 1. palindrome
	// result := codes.Palindrome("vikkiv")
	// println(result)

	// // 2. reverse no
	// result := codes.RevNum(12345)
	// println(result)

	// // 3. reverse no
	// codes.Fibonacci(12)

	// // 4. amstrong no
	// result := codes.AmstrongNum(153)
	// println(result)

	// // 5. prime no
	// result := codes.PrimeNum(11)
	// println(result)

	// // 6. perfect no
	// result := codes.PerfectNum(33550336)
	// println(result)

	// //***************************************matrixes operations***************************************

	// // initialising matrices
	// a := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// b := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	// // 1. add no
	// res := matrix.Add(a, b)

	// // 2. sub no
	// res := matrix.Sub(a, b)

	// // 3. mul no
	// res := matrix.Mul(a, b)

	// // printing result
	// for i := 0; i < 3; i++ {
	// 	for j := 0; j < 3; j++ {
	// 		print(res[i][j], ", ")
	// 	}
	// 	println("")
	// }

	// //***************************************patterns operations***************************************

	// // 1. left triangle
	// star.LeftTriangle(7)

	// // 2. right triangle
	// star.RightTriangle(7)

	// 3. heart
	star.Heart(15)
}
