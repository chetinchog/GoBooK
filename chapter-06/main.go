package main

import (
	"fmt"
	"time"
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
	ms := make(map[string]int)
	ms["key"] = 10
	fmt.Println(ms, ms["key"])
	mi := make(map[int]int)
	mi[1] = 10
	fmt.Println(mi, mi[1])

	elementsMake := make(map[string]string)
	elementsMake["H"] = "Hydrogen"
	elementsMake["He"] = "Helium"
	elementsMake["Li"] = "Lithium"
	fmt.Println(elementsMake)

	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}

	fmt.Println(elements)
	fmt.Println(elements["Li"])
	fmt.Println(elements["Un"])
	if name, ok := elements["Un"]; ok {
		fmt.Println(name, ok)
	} else {
		fmt.Printf("Element %s not found!\n", `"Un"`)
	}

	superElements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "solid",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": map[string]string{
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": map[string]string{
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": map[string]string{
			"name":  "Neon",
			"state": "gas",
		},
	}

	if el, ok := superElements["Li"]; ok {
		fmt.Println("| Li", "\t|", el["name"], "\t|", el["state"])
	}
	if el, ok := superElements["F"]; ok {
		fmt.Println("| F", "\t|", el["name"], "\t|", el["state"])
	}
	if el, ok := superElements["He"]; ok {
		fmt.Println("| He", "\t|", el["name"], "\t|", el["state"])
	}

	fmt.Println("- End -")
	// que := []int{
	// 	48, 96, 86, 68, 48, 96, 86, 68, 48, 96, 86, 68, 48, 96, 86, 68,
	// 	57, 82, 63, 70, 57, 82, 63, 70, 57, 82, 63, 70, 57, 82, 63, 70,
	// 	37, 34, 83, 27, 37, 34, 83, 27, 37, 34, 83, 27, 37, 34, 83, 27,
	// 	19, 97, 9, 17, 19, 97, 9, 17, 19, 97, 9, 17, 19, 97, 9, 17, 17,
	// }
	big := make(map[int]int)
	for i := range make([]int8, 100000000) {
		big[i] = i
	}
	fmt.Println("Searching min in", len(big)/1000000, "M items")
	// timer := time.Now().Second()
	fmt.Println("Time start:", time.Now().Second(), "s |", time.Now().Nanosecond()/100000000, "ms")
	small(big)
	fmt.Println("Time end:", time.Now().Second(), "s |", time.Now().Nanosecond()/100000000, "ms")
}

func small(x map[int]int) {
	s := x[0]
	for _, v := range x {
		if v < s {
			s = v
		}
	}
	fmt.Println("Min item:", s)
}
