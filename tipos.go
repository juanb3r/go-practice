package main

import (
	"fmt"
	"strconv"
)

func main() {
	age := 22
	quantity := "40"
	// Itoa() converts int to string
	ageStr := strconv.Itoa(age)

	fmt.Println("Mi edad es " + ageStr)
	// Atoi() converts string to int, this return to values
	quantityInt, _ := strconv.Atoi(quantity)
	fmt.Println("La cantidad es " + strconv.Itoa(quantityInt))
	fmt.Printf("La cantidad es %d", quantityInt)
}
