package repository

import "github.com/maccath/go-webapp/users"

var UserStore map[string]*users.User

type UserMemoryRepo struct {
}

func init() {
	UserStore = make(map[string]*users.User)
}

func (r *UserMemoryRepo) FindByName(name string) *users.User {
	return UserStore[name]
}

func (r *UserMemoryRepo) Save(user users.User) {
	UserStore[user.Name] = &user
}
