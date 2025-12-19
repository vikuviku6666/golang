package main

import "fmt"

type bot interface {
	getGreetings() string
}

type englishBot struct {}
type spanishBot struct {}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreetings(eb)
	printGreetings(sb)


}

func (englishBot) getGreetings() string {
	return "Hello"
}

func (spanishBot) getGreetings() string {
	return "Hålla"
}

func printGreetings(b bot) {
	fmt.Println(b.getGreetings())
}

