// Nos sirve para hacer la copia a los arreglos
package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4}
	copia := make([]int, 2)

	// Copy solo copia la cantidad de elementos de acuerdo a la capacidad
	// copia desde el segundo argumento al primero
	copy(copia, slice)

	fmt.Println(slice)
	fmt.Println(copia)
}
