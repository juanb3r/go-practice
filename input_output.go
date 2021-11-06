package main

import "fmt"

func main() {
	var age int
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &name)
	fmt.Printf("Hola %s Ingresa tu edad: ", name)
	fmt.Scanf("%d", &age)
	fmt.Printf("Tu edad es: %d\n", age)
}
