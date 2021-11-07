package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
Un problema de la lectura de archivos es que se puede haber interrumpido el programa
y no haber leido todo el archivo y menos cerrado. Para evitar eso, se puede usar
defer.
*/
func main() {
	ejecucion := leerArchivo()
	fmt.Println(ejecucion)
}

func leerArchivo() bool {
	file, err := os.Open("./reporte.csv")
	defer func() {
		file.Close()
		fmt.Println("Se cierra el archivo")
	}()
	if err != nil {
		fmt.Println(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			linea := scanner.Text()
			fmt.Println(linea)
		}
		return true
	}
	return false
}
