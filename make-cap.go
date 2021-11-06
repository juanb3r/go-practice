package main

import "fmt"

func main() {
	// Make sirve para crear un slice, el tercer parametro indica la capacidad
	// el append nos sirve para agregar otro valor al slice gracias al capacity
	capacity := 0
	fmt.Println("Ingrese capacidad:")
	fmt.Scanf("%d", &capacity)
	slice := make([]int, 5, capacity)
	fmt.Println(slice)
	fmt.Println(cap(slice))
	slice = append(slice, 8)
	fmt.Println(slice)
}
