package main

import "fmt"

// This means this function returns a string.
func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("world"))
}