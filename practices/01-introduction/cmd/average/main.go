package main

import "fmt"

func main() {
	a, b, c := 20, 25, 35
	sum := a + b + c
	var average float64 = float64(sum)
	average /= 3
	fmt.Printf("Среднее время: %.2f", average)
}
