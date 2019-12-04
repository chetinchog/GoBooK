package main

import (
	"fmt"
)

// Main function
func main() {
	fmt.Println("- Start -")

	fmt.Println("- Arrays -")
	var x [5]int
	x[4] = 100
	fmt.Println(x)

	var y [5]float64
	y[0] = 98
	y[1] = 93
	y[2] = 77
	y[3] = 82
	y[4] = 83

	var total float64 = 0
	for i := 0; i < len(y); i++ {
		total += y[i]
	}
	fmt.Println(total / float64(len(y)))

	z := [5]float64{98, 93, 77, 82, 83}
	total = 0
	for _, value := range z {
		total += value
	}
	fmt.Println(total / float64(len(y)))

	fmt.Println("- Slices -")
	var s []float64
	s = make([]float64, 5, 10)
	fmt.Println(s)
	arr := [5]float64{1, 2, 3, 4, 5}
	s = arr[0:5]
	fmt.Println(s)
	s = arr[0:]
	fmt.Println(s)
	s = arr[:5]
	fmt.Println(s)
	s = arr[:]
	fmt.Println(s)

	slice1 := []int{1, 2, 3}
	fmt.Println(slice1)
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice2)
	slice3 := make([]int, 2)
	copy(slice3, slice1)
	fmt.Println(slice3)

	fmt.Println("- Maps -")

	fmt.Println("- End -")
}
