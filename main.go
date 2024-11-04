package main

import (
    "fmt"
    "golangproject/utils"
)

func main() {
	n := 4 // Declare n here

	// Display results of utility functions
	fmt.Println("MagicSum:", utils.MagicSum(n))
	fmt.Println("MagicPow:", utils.MagicPow(n))
	fmt.Println("MagicOdd:", utils.MagicOdd(n))
	fmt.Println("MagicGrade:", utils.MagicGrade(n))
	fmt.Println("MagicName:", utils.MagicName(n))
	fmt.Println("MagicTria:", utils.MagicTria(n))

	// Use MagicChange and display result before and after
	x := n
	fmt.Println("Before MagicChange:", x)
	utils.MagicChange(&x)
	fmt.Println("After MagicChange:", x)

	// Initialize MagicNumber struct and use Multiply method
	magicNumber := utils.MagicNumber{Number: n}
	fmt.Println("Before Multiply:", magicNumber.Number)
	magicNumber.Multiply(n)
	fmt.Println("After Multiply:", magicNumber.Number)
}