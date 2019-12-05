package main

import (
	"fmt"
)

func average(xs []float64) (avg float64) {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	avg = total / float64(len(xs))
	return
}

func double() (x int, ok bool) {
	x = 1
	ok = true
	return
}

func add(args ...int) (total int) {
	for _, v := range args {
		total += v
	}
	return
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func fibonacci(x uint) uint {
	if x == 0 || x < 0 {
		return 0
	} else if x == 1 {
		return x
	}
	return fibonacci(x-1) + fibonacci(x-2)
}

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

func catch(foo func(interface{})) {
	foo(recover())
}

func mainRecover(err interface{}) {
	if err != nil {
		fmt.Println("Something went wrong! Reazon:", err)
	}
}

func main() {
	defer catch(mainRecover)

	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))
	if x, ok := double(); ok {
		fmt.Println(x)
	}
	fmt.Println(add(1, 2, 3, 4))
	xint := []int{3, 4}
	xint[1] = 3
	fmt.Println(add(xint...))

	sum := func(x, y int) int {
		return x + y
	}
	fmt.Println(sum(1, 1))

	y := 0
	increment := func() int {
		y++
		return y
	}
	fmt.Println(increment())
	fmt.Println(increment())

	nextEven := makeEvenGenerator()
	fmt.Println("nextEven", nextEven())
	fmt.Println("nextEven", nextEven())
	fmt.Println("nextEven", nextEven())

	fmt.Println("Factorial:", factorial(65))

	defer second()
	first()

	for index := 0; index <= 8; index++ {
		fmt.Printf("fibonacci(%d) -> %d\n", index, fibonacci(uint(index)))
	}

	panic("Panic!")
}
