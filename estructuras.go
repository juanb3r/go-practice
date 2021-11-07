package main

import "fmt"

// Declaracion de una estructura llamada User, se crea fuera de la función main
type User struct {
	nombre, apellido string
	edad             int
}

func main() {
	// Creamos un usurio vacio
	var salvador User
	// Se define un usuario
	var juanber = User{
		nombre:   "Juan",
		apellido: "Perez",
		edad:     25,
	}
	// con new nos devuelve la dirección del puntero de memoria
	andres := new(User)
	fmt.Println("Devuelve el puntero con el &")
	fmt.Println(andres)
	fmt.Println("Devuelve el valor de la variable")
	fmt.Println(*andres)
	// No hay necesidad de usar el apuntador *
	andres.nombre = "Andrés"
	fmt.Println(andres.nombre)
	fmt.Println(juanber)
	fmt.Println(salvador)
}
