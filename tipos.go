package main

import (
	"fmt"
	"strconv"
)

func main() {
	age := 22
	quantity := "40"
	// Itoa() converts int to string
	age_str := strconv.Itoa(age)

	fmt.Println("Mi edad es " + age_str)
	// Atoi() converts string to int, this return to values
	quantity_int, _ := strconv.Atoi(quantity)
	fmt.Println("La cantidad es " + strconv.Itoa(quantity_int))
	fmt.Printf("La cantidad es %d", quantity_int)
}
