package main

import "fmt"

// This means this function returns a string.
func Hello() string {
	return "Hello, World!"
}

func main() {
	fmt.Println(Hello())
}