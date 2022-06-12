package memory

import (
	"github.com/maccath/go-webapp/pkg/user"
)

type Repository map[string]*user.Model

func (r Repository) FindByName(name string) *user.Model {
	return r[name]
}

func (r Repository) Save(user user.Model) {
	r[user.Name] = &user
}
