package memory

import (
	"testing"
)

func TestRepository_FindByName(t *testing.T) {
	var repo = Repository{}
	repo.Save(User{"katy"})

	t.Run("It finds a known user", func(t *testing.T) {
		var user = repo.FindByName("katy")

		var got = user.Name
		var want = "katy"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})

	t.Run("It doesn't find an unknown user", func(t *testing.T) {
		var user = repo.FindByName("alex")

		var got = user.Name
		var want = ""

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})
}

func TestRepository_Save(t *testing.T) {
	var repo = Repository{}

	t.Run("A user doesn't exist", func(t *testing.T) {
		var user = repo.FindByName("alex")

		var got = user.Name
		var want = ""

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})

	t.Run("It saves a given user", func(t *testing.T) {
		repo.Save(User{"alex"})

		var user = repo.FindByName("alex")

		var got = user.Name
		var want = "alex"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})

	t.Run("The saved user exists", func(t *testing.T) {
		var user = repo.FindByName("alex")

		var got = user.Name
		var want = "alex"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})
}
