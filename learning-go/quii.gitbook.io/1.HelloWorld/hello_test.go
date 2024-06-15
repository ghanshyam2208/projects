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
		got := Hello(name)
		want := englishHelloPrefix + name

		if got != want {
			assertCorrectMessage(t, got, want)
		}
	})

	t.Run("Hello world when empty string is passed", func(t *testing.T) {
		got := Hello("")
		want := englishHelloPrefix + "world"

		if got != want {
			assertCorrectMessage(t, got, want)
		}
	})
}
