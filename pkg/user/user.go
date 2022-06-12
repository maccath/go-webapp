package user

type Model struct {
	Name string
}

type Repository interface {
	Save(u Model)
	FindByName(name string) *Model
}
