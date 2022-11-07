package main

import "fmt"

func main() {
	num1 := 5
	num2 := 5
	if num1 < num2 {
		fmt.Printf("If condition is true")
	} else if num1 > num2 {
		fmt.Printf("Else if condition is true")
	} else {
		fmt.Printf("nothing above is true")
	}

}
