package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Scanf("%d %d", &num1, &num2)
	f := float64(num1) / float64(num2)
	fmt.Printf("%.9f", f)

}
