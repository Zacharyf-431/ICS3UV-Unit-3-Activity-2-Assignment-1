/*
	* Author: Zachary Fowler
	* Version: 1.0.0
	* Date: 2025-11-17
	* This file rounds given numbers
	*/

package main

import "fmt"

func main() {
	// assign constants 
	const num1 float64 = 8.5467
	const num2 float64 = 9.6382
	const num3 float64 = 18.5146
	const num4 float64 = 125.496

	// INPUT - No input

	// OUTPUT
	// rounded values with field widths
	fmt.Printf("%10.3f\n", num1)
	fmt.Printf("%8.5f\n", num2)
	fmt.Printf("%6.1f\n", num3)
	fmt.Printf("%3.1f\n", num4)

	fmt.Println("\nDone.")
	}