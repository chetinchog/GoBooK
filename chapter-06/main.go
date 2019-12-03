package main

import (
	"fmt"
	"time"
)

// Main function
func main() {
	fmt.Println("- Start -")
	timer := time.Now().Second()
	for index := 0; index < 28000000000; index++ {
		if index%1000000000 == 0 {
			fmt.Println(time.Now().Second(), " - ", time.Now().Second()-timer)
			timer = time.Now().Second()
		}
	}
	fmt.Println("- End -")
}
