package main

import "fmt"


const englishHelloPrefix = "Hello, "

// This means this function accept a string and returns a string.
func Hello(name string, language string) string {

	if name == "" {
		name = "World"
	}

	if language == "Spanish" {
		return "Hola, " + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}