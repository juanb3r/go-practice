package main

import "fmt"

func main() {
	/*
		Los punteros son una referencia a una dirección de memoria.
		Para acceder a una dirección de memoria, se utiliza el operador &
		X,Y => ad342f => 3
		X => 4
		Y => 4

		*T => *int,  *string, *Struct
		valor zero => nil
	*/
	var x, y *int
	entero := 5
	// Obtenemos la dirección del espacio de la memoria con el puntero &
	x = &entero
	y = &entero
	// Cambiamos el valor del espacio guardado en la memoria con *
	*y = 14
	// Imprimimos el valor guardado en el espacio de memoria y la dirección
	fmt.Printf("%d %d", *x, y)
}
