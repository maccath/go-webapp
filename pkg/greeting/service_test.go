package greeting

import (
	"testing"
)

type mockRepository struct{}

// mockRepository has only one user - katy.
func (repo mockRepository) FindByName(name string) User {
	if name == "katy" {
		return User{name}
	}

	return User{}
}

func TestGreet(t *testing.T) {
	var vars = make(map[string]string)
	var repo = mockRepository{}
	var service = NewService(repo)

	t.Run("It greets the world when no user name is specified", func(t *testing.T) {
		var got = service.Greet(vars["name"])
		var want = "Hello, world!"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})

	t.Run("It greets the user name when an known user's name is specified", func(t *testing.T) {
		vars["name"] = "katy"

		var got = service.Greet(vars["name"])
		var want = "Hello, katy!"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})

	t.Run("It greets the world when an unknown user's name is specified", func(t *testing.T) {
		vars["name"] = "alex"

		var got = service.Greet(vars["name"])
		var want = "Hello, world!"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})
}
