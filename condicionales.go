package main

import "fmt"

/*
	== igual
	!= diferente
	> mayor que
	< menor que
	>= mayor o igual que
	<= menor o igual que
	&& y
	|| o
*/
func main() {
	x := 10
	y := 20
	if x < y {
		fmt.Printf("%d es mayor que %d\n", y, x)
	} else {
		fmt.Printf("%d es menor que %d\n", y, x)
	}
}
