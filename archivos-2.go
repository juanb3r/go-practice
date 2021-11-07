package main

import (
	"bufio"
	"fmt"
	"os"
)

// Leyendo archivo linea por linea
func main() {
	file, err := os.Open("./reporte.csv")
	if err != nil {
		fmt.Println(err)
	} else {
		scanner := bufio.NewScanner(file)

		var i int
		for scanner.Scan() {
			i++
			linea := scanner.Text()
			fmt.Println(i)
			fmt.Println(linea)
		}
	}
}
