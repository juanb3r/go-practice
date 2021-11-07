package main

import "fmt"

/*
Las interfaces son un tipo de datos que obligan a las
estructuras a implementar sus métodos.
*/

// Los nombres de los metodos en minusculas inicial son privados
type User interface {
	permisos() int // El método permisos() es obligatorio
	obtenerNombre() string
}

// Se crea la estructura pero no se hace explicitamente la implementacion
// de la interfaz
type Admin struct {
	name string
}

type Editor struct {
	name string
}

// Pero si añadimos el método permisos() a la estructura Admin,
// podemos estamos implementando la interfaz.
func (this Admin) permisos() int {
	return 1
}

func (this Admin) obtenerNombre() string {
	return this.name
}

func (this Editor) permisos() int {
	return 2
}

func (this Editor) obtenerNombre() string {
	return this.name
}

func auth(user User) bool {
	if user.permisos() == 1 {
		return true
	}
	return false
}

func main() {
	// admin := Admin{"Juanb3r"}
	editor := Editor{"Peter Parker"}
	if auth(editor) {
		fmt.Println("El usuario tiene permisos")
	} else {
		fmt.Println("El usuario no tiene permisos")
	}
}
