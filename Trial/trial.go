package main
 import (
	"fmt"
 )

 func main(){
	var name string 
	var age int
	fmt.Println("Please enter your name: ")
	fmt.Println("Please enter your age: ")
	fmt.Scanln(&name)
	fmt.Scanln(&age)

	fmt.Println("Hello:", name)
	fmt.Println("I am ", age , " years old.")

	if (name == ""){
fmt.Println("Name must be inputted")
	}

	var c float32
	var d float32
	var operator string


	fmt.Println("Enter first number: ")
	fmt.Scanln(&c)
	fmt.Println("Enter second number: ")
	fmt.Scanln(&d)

	fmt.Println("Enter operator:(+, -, *, /) ")
	fmt.Scanln(&operator)

	switch operator{
	case "+":
		fmt.Println("Result: ", add(c, d))
	case "-":
		fmt.Println("Result: ", subtract(c,d))	
	case "*":
		fmt.Println("Result: ", multiply(c,d))
	case "/":
		fmt.Println("Result: ", divide(c,d))		
	}

 }
 func add(a, b float32)float32{
	return a + b
 }
 func subtract(a,b float32)float32{
	return a - b
 }
 func multiply(a, b float32)float32{
	return a * b
 }
 func divide(a,b float32)float32{
	if b == 0{
		fmt.Println("Cannot divide")
	}
	return a/b
 }