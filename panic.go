package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
Panic es una función que se ejecuta para levantar un error.
se puede detener con el metodo recover()
*/
func main() {
	ejecutarLeerArchivo()
	fmt.Println("Creo que me imprimiré gracias a Recover")
}

func ejecutarLeerArchivo() {
	leerArchivo()
	fmt.Println("Ejecutando leerArchivo")
}

func leerArchivo() bool {
	file, err := os.Open("./reporte.csv")
	defer func() {
		file.Close()
		fmt.Println("Se cierra el archivo")

		recover()
	}()
	if err != nil {
		panic(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			linea := scanner.Text()
			fmt.Println(linea)
		}
		return true
	}
}
