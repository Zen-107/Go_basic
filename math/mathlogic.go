package main
import "fmt"

func main(){
	Num1 := 12
	Num2 := 5
	var Num3 float32 = float32(Num1) / float32(Num2)
	var Num4 float32 = float32(Num2)
	
	fmt.Println("What is 12+5??? ",Num1 + Num2)
	fmt.Println("What is 12-5??? ",Num1 - Num2)
	fmt.Println("What is 12*5??? ",Num1 * Num2)
	fmt.Println("What is (12/5)/Num2??? ",Num3 / Num4)
	fmt.Println("What is 12%5??? ",Num1 % Num2)
}