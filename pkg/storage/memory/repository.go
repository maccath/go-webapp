package memory

import (
	"github.com/maccath/go-webapp/pkg/greeting"
)

type Repository map[string]User

func (r Repository) FindByName(name string) greeting.User {
	var user greeting.User

	if u, ok := r[name]; ok {
		user.Name = u.Name
	}

	return user
}

func (r Repository) Save(user User) {
	var u User

	u.Name = user.Name

	r[user.Name] = u
}
