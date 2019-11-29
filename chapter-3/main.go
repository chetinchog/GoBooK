package main

import (
	"fmt"
)

func main() {
	fmt.Println("\n> Numbers")

	fmt.Println("\n> > Integers")
	fmt.Println("1 + 1 =", 1+1)

	fmt.Println("\n> - Floating Point Numbers")
	fmt.Println("1 + 1 =", 1.0+1.0)

	fmt.Println("\n> Strings")
	fmt.Println(len("Hello World!"))
	fmt.Println("Hello World!"[1])
	fmt.Println("Hello " + "World!")

	fmt.Println("\n> Booleans")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

	fmt.Println("\n> Multiplication")
	fmt.Println("32132 × 42452 =", 32132*42452)

	fmt.Println("\n> Boolean operation")
	fmt.Println("(true && false) || (false && true) || !(false && false) =", (true && false) || (false && true) || !(false && false))
}
