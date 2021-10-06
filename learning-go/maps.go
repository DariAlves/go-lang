package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {
	myMap := make(map[string]User)

	log.Println(myMap)

	user1 := User{
		FirstName: "Jack",
		LastName:  "Sparrow",
	}

	myMap["user1"] = user1

	log.Println(myMap)
	log.Println(user1)
	log.Println(myMap["user1"])
	log.Println(myMap["user1"].FirstName)
}

// func main() {
// 	myMap := make(map[string]string)

// 	myMap["dog"] = "bob"
// 	myMap["cat"] = "tom"
// 	myMap["bird"] = "blue"

// 	log.Println(myMap)
// 	log.Println(myMap["dog"])
// }

// func main() {
// 	myMap := make(map[string]int)

// 	myMap["one"] = 1
// 	myMap["two"] = 2

// 	log.Println(myMap)
// 	log.Println(myMap["two"])
// }
