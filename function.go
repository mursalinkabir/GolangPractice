package main

import "fmt"

func main() {
	var a int = 100
	var b int = 200
	var ret int 

	ret = max(a,b)
	fmt.Printf(" The max value is %d\n", ret)
}

func max(num1,num2 int) int {
	var res int
	if num1>num2 {
		res = num1
	}else{
		res = num2
	}
	return res
}