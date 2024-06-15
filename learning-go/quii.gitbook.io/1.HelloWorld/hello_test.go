package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got string, want string) {
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		const name = "krishna"
		got := Hello(name, "English")
		want := englishHelloPrefix + name

		assertCorrectMessage(t, got, want)
	})

	t.Run("Hello world when empty string is passed", func(t *testing.T) {
		got := Hello("", "English")
		want := englishHelloPrefix + "world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("greet in spanish ", func(t *testing.T) {
		const name = "Elodie"
		got := Hello(name, "Spanish")
		want := "Hola, " + name
		assertCorrectMessage(t, got, want)
	})

	t.Run("greet in french ", func(t *testing.T) {
		const name = "Elodie"
		got := Hello(name, "French")
		want := "Bonjour, " + name
		assertCorrectMessage(t, got, want)
	})
}
