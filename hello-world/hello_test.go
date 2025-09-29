package main

import "testing"

// t of type *testing.T is our "hook" into the testing framework
func TestHello(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {
		// Declaring variables
		got := Hello("Chris")
		want := "Hello, Chris"

		// Calling the Errorf method on our t, which will print 
		// out a message and fail the test. The f stands for format,
		// which allows us to build a string with values inserted into 
		// the placeholder values %q. When you make the test fail, 
		// it should be clear how it works.
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}