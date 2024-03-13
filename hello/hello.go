package main

import "fmt"

const (
	spanish = "Spanish"
	french = "French"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Salut, "
)

// Hello returns a greeting for the named person.
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	if language == spanish {
		return spanishHelloPrefix + name + "."
	}
	if language == french {
		return frenchHelloPrefix + name + "."
	}
	return englishHelloPrefix + name + "."
}

func main() {
	name := "Pello"
	fmt.Println(Hello(name, ""))
}
