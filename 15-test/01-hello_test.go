package main

import (
	"testing"
)

func TestHelloInitialVersion(t *testing.T) {
	expected := "Hello World!"
	actual := Hello("World!", "")

	if actual != expected {
		t.Errorf("got '%s' - wanted '%s'", actual, expected)
	}

}

func TestHelloYou(t *testing.T) {
	expected := "Hello Javi"
	actual := Hello("Javi", "")

	if actual != expected {
		t.Errorf("got '%s' - wanted '%s'", actual, expected)
	}
}

func TestHelloEmpty(t *testing.T) {

	t.Run("say  'Hello World!' when an empty string is supplied", func(t *testing.T) {
		expected := "Hello World!"
		actual := Hello("", "")

		if actual != expected {
			t.Errorf("got '%s' - wanted '%s'", actual, expected)
		}
	})
}

//tests refactor

func TestHelloRefactored(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, actual string, expected string) {
		t.Helper() //when it fails, will report the function call rather than inside this
		if actual != expected {
			t.Errorf("got '%s' - wanted '%s'", actual, expected)
		}
	}

	t.Run("should default 'Hello World!' when empty name", func(t *testing.T) {
		expected := "Hello World!"
		actual := Hello("", "")
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("say hello to whatever name", func(t *testing.T) {
		expected := "Hello Javi!"
		actual := Hello("Javi!", "")
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("say hello in Spanish", func(t *testing.T) {
		expected := "Hola Javi!"
		actual := Hello("Javi!", "Spanish")
		assertCorrectMessage(t, actual, expected)
	})
}
