package main

import "fmt"

//Definimos dos estructuras

type Transporte struct {
	capacidadPasajeros int
	ruedas             bool
}

type Terrestre struct {
	Transporte
}

type Automovil struct {
	Terrestre
	color           string
	llantas         int
	velocidadMaxima int
}

func (this Transporte) mover() {
	fmt.Println("Moviendose")
}

func (this Terrestre) mover() {
	fmt.Println("Moviendose en el terreno")
}

func (this Automovil) acelerar(
	vi float32,
	vf float32,
	ti float32,
	tf float32) float32 {
	return (vf - vi) / (tf - ti)
}

func main() {
	ferrari := Automovil{
		Terrestre{
			Transporte{2, true}},
		"rojo", 4, 326}
	println(ferrari.color)
	ferrari.mover()
	fmt.Printf("Acelera a %f km/s2\n", ferrari.acelerar(0, 100, 0, 7))
}
