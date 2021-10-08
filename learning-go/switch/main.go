package main

import "log"

func main() {
	myVar := "cat"

	switch myVar {
	case "cat":
		log.Println("myVar is set to cat")

	case "dog":
		log.Println("myVar is set to dog")

	case "bird":
		log.Println("myVar is set to bird")

	default:
		log.Println("myVar is set to default value")
	}
}
