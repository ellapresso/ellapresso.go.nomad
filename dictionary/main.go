package main

import (
	"fmt"

	"github.com/ellapresso/learngo/dict"
)

func main() {
	dictionary := dict.Dictionary{}
	// dictionary["hello"] = "안녕"
	// fmt.Println(dictionary)
	// dictionary := dict.Dictionary{"first": "First word"}
	// fmt.Println(dictionary["first"])
	// fmt.Println(dictionary.Search("first"))
	// fmt.Println(dictionary.Search("second"))

	// definition, err := dictionary.Search("first")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(definition)
	// }

	// dictionary.Add("hello", "Greeting")
	// definition, err = dictionary.Search("hello")

	//
	// err := dictionary.Add("hello", "Greeting")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// hello, err := dictionary.Search("hello")
	// fmt.Println(hello)
	// err2 := dictionary.Add("hello", "Greeting2")
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }

	//
	dictionary.Add("goodbye", "Bye")
	dictionary.Add("gohome", "Go Home")
	err := dictionary.Update("gohomessss", "yes")
	if err != nil {
		fmt.Println(err)
	}

	err2 := dictionary.Delete("gohomesss")
	if err2 != nil {
		fmt.Println(err2)
	}
}
