package main

import "fmt"

func main() {
	var name string = "Zen"
	var age int = 25
	var isStudent bool = true
	const mass float32 = 50.00
	hight := 100

	fmt.Println("My hight is ", hight, "cm.")

	hight = 105
	fmt.Println("My hight is ", hight, "cm.")
	fmt.Println("")
	fmt.Println("My name is ", name)
	fmt.Println("My age is ", age)
	fmt.Println("My mass is ", mass, "kg.")

	if isStudent {
		fmt.Println("I am a student")
	} else {
		fmt.Println("I am not a student")
	}

	//How to show data type
	fmt.Println()
	fmt.Printf("Data type of name:\t\t%T\n", name)
	fmt.Printf("Data type of age:\t\t%T\n", age)
	fmt.Printf("Data type of isStudent:\t\t%T\n", isStudent)
	fmt.Printf("Data type of mass:\t\t%T\n", mass)
	fmt.Printf("Data type of hight:\t\t%T\n", hight)
}
