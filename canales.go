package main

import "fmt"

func main() {
	channel := make(chan string)

	go func(channel chan string) {
		for {
			var name string
			fmt.Scanln(&name)
			channel <- name
		}
	}(channel)

	// recibiendo datos del canal
	msg := <-channel
	fmt.Println("Hello for first time", msg)

	msg = <-channel
	fmt.Println("Hello again", msg)
	/*
		la go rutina se detiene hasta que el canal recibe dos veces
	*/
}
