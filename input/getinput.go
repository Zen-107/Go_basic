package main

import "fmt"

func main() {
	var unitcm float64
	fmt.Println("Enter unit in cm:")
	fmt.Scanf("%f", &unitcm)
	fmt.Printf("Input is %.2f\n", unitcm)
	const cmtofeet = 0.0328
	unitfeet := unitcm * cmtofeet
	fmt.Println("Convert to feet is ",unitfeet)
}
