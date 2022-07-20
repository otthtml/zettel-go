package main

import "fmt"

func main() {
	eb := englishBot{}
	printGreeting(eb)
	sb := spanishBot{}
	printGreeting(sb)
}

type englishBot struct{}

type spanishBot struct{}

func (englishBot) getGreeting() string {
	return "hello there"
}

func (spanishBot) getGreeting() string {
	return "hola"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

type bot interface {
	getGreeting() string
}
