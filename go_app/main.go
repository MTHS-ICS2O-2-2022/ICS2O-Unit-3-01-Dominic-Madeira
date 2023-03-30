// Created by: Dominic M.
// Created on: March 2023
//
// This program finds the area of a trapezoid.
package main

import "fmt"

func main() {
	var aBase float64
	var bBase float64
	var height float64
	var area float64

	// input
	fmt.Println("This program finds the area of a trapezoid.")
	fmt.Println()
	fmt.Print("Enter a length for base a: ")
	fmt.Scanln(&aBase)
	fmt.Println()
	fmt.Print("Enter a length for base b: ")
	fmt.Scanln(&bBase)
	fmt.Println()
	fmt.Print("Enter a height: ")
	fmt.Scanln(&height)

	// process
	area = ((aBase + bBase) / 2) * height

	// output
	fmt.Println()
	fmt.Println("The area is:", area, "cm²")

	fmt.Println("\nDone.")
}
