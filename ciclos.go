package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// el ciclo for con esa condicional se comporta como el while
	j := 0
	for j < 10 {
		fmt.Println(j)
		j++
	}
}
