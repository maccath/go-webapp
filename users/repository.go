package users

type Repository interface {
	Save(u User)
	FindByName(name string) *User
}
