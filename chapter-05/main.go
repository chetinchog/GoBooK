package main

import (
	"fmt"
)

// Main function
func main() {
	fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)
	fmt.Println(4)
	fmt.Println(5)
	fmt.Println(6)
	fmt.Println(7)
	fmt.Println(8)
	fmt.Println(9)
	fmt.Println(10)
	fmt.Println("-")
	fmt.Println(`1
	2
	3
	4
	5
	6
	7
	8
	9
	10`)
	fmt.Println("-")
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}
	fmt.Println("-")

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
		switch i {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		case 4:
			fmt.Println("Four")
		case 5:
			fmt.Println("Five")
		default:
			fmt.Println("Unknown Number")
		}
	}

	fmt.Println("-")
	h := 10
	if h > 10 {
		fmt.Println("Big")
	} else {
		fmt.Println("Small")
	}

	fmt.Println("-")
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}

	fmt.Println("-")
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

	fmt.Println("- Start -")
	for i := 1; i <= 10000000000; i++ {
		if i%1000000000 == 0 {
			fmt.Println(i)
		}
	}
	fmt.Println("- End -")
}
