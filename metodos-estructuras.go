/*
En esta practica crearemos métodos para las estructuras
*/
package main

import "fmt"

type User struct {
	nombre, apellido string
	edad             int
}

/* Creamos una función que recibe como parametro un dato de tipo User
y retorna un dato de tipo string que contiene el nombre y el apellido del
usuario
*/
func (usuario User) nombre_completo() string {
	return usuario.nombre + " " + usuario.apellido
}

/* Cuando se crea un metodo para modificar la estructura, esta crea una copia
si no lleva el apuntador (*)
*/
func (usuario *User) set_nombre(nombre string) {
	usuario.nombre = nombre
}

func main() {
	var salvador User
	salvador.nombre = "Salvador"
	salvador.apellido = "Bermudez"
	salvador.set_nombre("Salvador Miguel")
	fmt.Println(salvador.nombre_completo())
}
