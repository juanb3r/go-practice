package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	dat, err := ioutil.ReadFile("./reporte.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", dat)
}
