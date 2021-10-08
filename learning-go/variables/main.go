package main

import "log"

// Global scope
var s = "one"

func main() {
	var s2 = "two"

	log.Println("s is", s)
	log.Println("s2 is", s2)

	saySomething("xxx")

}

func saySomething(s3 string) (string, string) {
	log.Println("s from the saySomething func is", s)
	return s3, "four"
}
