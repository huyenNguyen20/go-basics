package main

import "fmt"

type englishBot struct {}
type spanishBot struct {}

func main() {
	eb := englishBot{}
	// sb := spanishBot{}

	printGreeting(eb)
	// printGreeting(sb)
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func (eb englishBot) getGreeting() string {
	// Codes to generate english readings
	return "Hi there!"
}

func (eb spanishBot) getGreeting() string {
	// Codes to generate spanish readings
	return "Hola"
}