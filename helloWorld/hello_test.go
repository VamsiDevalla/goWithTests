package helloworld

import "testing"

func TestHello(t *testing.T) {
	assertGreetingsMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("want %q but got %q", want, got)
		}
	}

	t.Run("Say hello to people", func(t *testing.T) {
		got := Hello("Test", "")
		want := "Hello, Test"
		assertGreetingsMessage(t, got, want)
	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertGreetingsMessage(t, got, want)
	})

	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Test", "spanish")
		want := "Hola, Test"
		assertGreetingsMessage(t, got, want)
	})

}
