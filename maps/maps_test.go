package main

import "testing"

func TestSearch(t *testing.T) {

	dict := Dictionary{"value": "this is a value"}

	t.Run("known word", func(t *testing.T) {

		got, _ := dict.Search("value")
		want := "this is a value"

		assertString(t, got, want)

	})

	t.Run("unknown word", func(t *testing.T) {

		_, err := dict.Search("unknown")

		if err == nil {
			t.Fatal("exected to get an error")
		}

		assertError(t, err, NotFoundErr)

	})

}

func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {

		dictionary := Dictionary{}

		err := dictionary.Add("test", "just a test")

		assertError(t, err, nil)

	})

	t.Run("existing word", func(t *testing.T) {

		dictionary := Dictionary{"test": "just a test"}

		err := dictionary.Add("test", "new test")

		assertError(t, err, KeyExistsErr)

	})

}

func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("want %s but got %s", want, got)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("want %q but got %q", want, got)
	}
}
