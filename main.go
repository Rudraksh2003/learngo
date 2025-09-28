package main

import "fmt"

func thing(num1 int, num2 int) {
	for i := num2; i <= num1; i++ {
		fmt.Println(i)
	}
}

func main() {
	fmt.Println("Lets learn Go")

	var num1, num2, choice int

	fmt.Print("Enter first num: ")
	fmt.Scan(&num1)

	fmt.Print("Enter second num: ")
	fmt.Scan(&num2)

	fmt.Print("Enter choice (1-Add, 2-Sub, 3-Mul, 4-Div): ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Println("Addition is", num1+num2)
	case 2:
		fmt.Println("Subtraction is", num1-num2)
	case 3:
		fmt.Println("Multiplication is", num1*num2)
	case 4:
		if num2 != 0 {
			fmt.Println("Division is", num1/num2)
		} else {
			fmt.Println("Error: Division by zero")
		}
	default:
		fmt.Println("Invalid choice")
	}

	if num1 > num2 {
		fmt.Println("Since num1 > num2, printing numbers in range:")
		thing(num1, num2)
	}
}
