package main

import "fmt"

const spanish = "Spanish"
const french = "French"	
const swedish = "Swedish"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const swedishHelloPrefix = "Hej, "

// This means this function accept a string and returns a string.
func Hello(name string, language string) string {

	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case swedish:
		prefix = swedishHelloPrefix
	}

	return prefix + name	
}

func main() {
	fmt.Println(Hello("world", ""))
}