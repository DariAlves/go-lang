package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {
	var mySlice []User

	user1 := User{
		FirstName: "Bruce",
		LastName:  "Wayne",
	}

	user2 := User{
		FirstName: "Clark",
		LastName:  "Kent",
	}

	user3 := User{
		FirstName: "Barry",
		LastName:  "Allen",
	}

	mySlice = append(mySlice, user1)
	mySlice = append(mySlice, user2)
	mySlice = append(mySlice, user3)

	for i, v := range mySlice {
		log.Println(i, v)
	}

}

// Map

// func main() {
// 	myMap := make(map[string]string)
// 	myMap["dog"] = "dog"
// 	myMap["cat"] = "cat"
// 	myMap["bird"] = "bird"

// 	for i, v := range myMap {
// 		log.Println(i, v)
// 	}
// }

// Slice

// func main() {
// 	animals := []string{"dog", "cat", "fish", "bird", "dragon"}

// 	for _, x := range animals {
// 		log.Println(x)
// 	}
// }

// For loops

// func main() {
// 	for i := 0; i <= 10; i++ {
// 		log.Println(i)
// 	}
// }
