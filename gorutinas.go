package main

import (
	"fmt"
	"time"
)

func main() {
	miNombreLentamente("Juan")
	fmt.Println("\nToca esperar es un poco de tiempo")
}

func miNombreLentamente(nombre string) {
	// letras := strings.Split(nombre, "") usa la libreria strings
	// usando la funcion split para separar las letras del nombre
	for _, letra := range nombre {
		fmt.Printf("%c", letra)
		time.Sleep(2 * time.Second)
	}
}
