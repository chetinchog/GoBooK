package main

import (
	"fmt"
	"log"
)

func zeroValue(x int) {
	x = 0
}

func zeroPointer(x *int) {
	*x = 0
}

func square(x *float64) {
	*x = *x * *x
}

func swap(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	log.Println("|\t", "C8!")

	x := 5
	zeroValue(x)
	fmt.Println(x)

	zeroPointer(&x)
	fmt.Println(x)

	fmt.Println("-")
	y := new(int)
	fmt.Println(*y)
	zeroPointer(y)
	fmt.Println(*y)

	z := 1.5
	square(&z)
	fmt.Println(z)

	fmt.Println("-")
	p := 1
	q := 2
	swap(&p, &q)
	fmt.Println(p, q)
}
