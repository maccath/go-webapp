package repository

import "github.com/maccath/go-webapp/users"

type UserMemoryRepo map[string]*users.User

func (r UserMemoryRepo) FindByName(name string) *users.User {
	return r[name]
}

func (r UserMemoryRepo) Save(user users.User) {
	r[user.Name] = &user
}
