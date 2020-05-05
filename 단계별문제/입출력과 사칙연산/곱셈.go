package main

import "fmt"

func main() {

	var num1, num2 int
	fmt.Scanf("%d", &num1)
	fmt.Scanf("%d", &num2)
	a := num2 % 10
	b := (num2 % 100) / 10
	c := (num2 / 100)
	fmt.Println(num1 * a)
	fmt.Println(num1 * b)
	fmt.Println(num1 * c)
	fmt.Println(num1 * num2)

}
