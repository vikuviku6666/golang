package main

import "fmt"

func main() {
	var names map[string]string
	rooms := make(map[string]string)
	fruits := map[string]string{
		"Sweet": "Apples",
		"Sour": "Limes",
	}
	names = make(map[string]string)
	names["long"] = "Vinay"
	rooms["large"] = "open"
	printMap(fruits)
	delete(fruits, "Sweet")


	fmt.Println(fruits)
	fmt.Println(rooms)
	fmt.Println(names)
}

func printMap(c map[string]string) {
	for fruit, value := range c {
		fmt.Println("The name of the ", fruit, "is", value)
	} 
}
