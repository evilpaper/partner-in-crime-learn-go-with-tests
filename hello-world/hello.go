package main

import "fmt"

// This means this function returns a string.
func Hello(name string) string {
	return "Hello, World!"
}

func main() {
	fmt.Println(Hello("world"))
}