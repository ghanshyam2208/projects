package dictionary

import (
	"testing"
)

func TestSearch(t *testing.T) {
	assertStrings := func(t testing.TB, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("test search works", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		actual, _ := Search(dictionary, "test")
		expected := "this is just a test"
		assertStrings(t, actual, expected)
	})

	t.Run("test for empty search", func(t *testing.T) {
		dictionary := Dictionary{"test": ""}

		actual, err := Search(dictionary, "unknown")
		expected := ""

		if err == nil {
			t.Error("this should shoot error")
		}

		if err != ErrNotFound {
			t.Error("Not found error did not match")
		}

		assertStrings(t, actual, expected)
	})

	t.Run("test dictionary addition ", func(t *testing.T) {
		dict := Dictionary{}

		dict.Add("test", "this is just a test")
		actual, _ := Search(dict, "test")
		expected := "this is just a test"
		assertStrings(t, actual, expected)
	})

	t.Run("test dictionary addition if word already exists ", func(t *testing.T) {
		dict := Dictionary{
			"test": "some value exists",
		}

		err := dict.Add("test", "this is just a test")
		if err == nil {
			t.Error("this should shoot error")
		}

		if err != ErrWordAlreadyExists {
			t.Error("word exists error did not match")
		}
	})
}
