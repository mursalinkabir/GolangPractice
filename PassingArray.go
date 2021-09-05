package main

import "fmt"

func main() {

	var balance = []int{100, 5, 10, 55, 16}
	var avg float32

	avg = getAverage(balance, 5)

	fmt.Printf("average value is: %f", avg)

}

func getAverage(arr []int, size int) float32 {
	var i, sum int
	var avg float32
	for i = 0; i < size; i++ {
		sum += arr[i]
	}
	avg = float32(sum / size)
	return avg
}
