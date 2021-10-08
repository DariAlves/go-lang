package main

import "log"

func main() {
	myNumber := 100
	isTrue := true

	if myNumber > 99 && !isTrue {
		log.Println("myNum is greater than 99 and isTrue is set to true")
	} else if myNumber < 100 && isTrue {
		log.Println("1")
	} else if myNumber > 101 && isTrue {
		log.Println("2")
	} else {
		log.Println("3")
	}
}

// func main() {
// 	var isTrue bool

// 	isTrue = false

// 	if isTrue == true {
// 		log.Println("isTrue is", isTrue)
// 	} else {
// 		log.Println("isTrue is", isTrue)
// 	}
// }
