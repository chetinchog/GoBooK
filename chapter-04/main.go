package main

import (
	"fmt"
)

var gx string = "Hello World"

func main() {
	var x string
	x = "first"
	fmt.Println(x)
	x = "second"
	fmt.Println(x)
	x = "first "
	fmt.Println(x)
	x = x + "second"
	fmt.Println(x)
	x = "hello"
	var y string = "world"
	fmt.Println(x == y)
	y = "hello"
	fmt.Println(x == y)
	z := 5 // Generally you should use this shorter form whenever possible.
	fmt.Println(z)
	fmt.Println(gx)
	f()

	const cx string = "HI CONST!"
	fmt.Println(cx)

	var (
		a = 5
		b = 10
		c = 15
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	input()
	feetToMeters()
	celcToFahr()
}

func f() {
	fmt.Println(gx)
}

func input() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)
}

func feetToMeters() {
	fmt.Print("Enter feets: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 0.3048
	fmt.Println(output, "mts")
}

func celcToFahr() {
	fmt.Print("Enter Celsius: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := (input / (5.0 / 9.0)) + 32.0
	fmt.Println(output, "Fahrenheit")
}
