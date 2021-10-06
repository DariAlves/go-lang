package main

import (
	"log"
	"sort"
)

func main() {
	var saiyans []string
	numbers := []int{10, 1, 8, 11, 6}

	saiyans = append(saiyans, "goku")
	saiyans = append(saiyans, "vegeta")
	saiyans = append(saiyans, "trunks")
	saiyans = append(saiyans, "gohan")

	log.Println(saiyans)
	sort.Strings(saiyans)

	log.Println(saiyans)
	log.Println(len(saiyans))
	log.Println(saiyans[0:3])

	log.Println(numbers)
	sort.Ints(numbers)
	log.Println(numbers)
}
