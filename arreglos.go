package main

import "fmt"

func main() {
	fmt.Println("Ingrese palabra:")
	var palabra string
	fmt.Scanf("%s", &palabra)
	largo := len(palabra)
	arreglo := [3]int{5, 8, 6}
	fmt.Println(arreglo)
	fmt.Println(largo)
	for i := 0; i < len(arreglo); i++ {
		fmt.Println(arreglo[i])
	}
}
