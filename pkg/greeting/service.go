package greeting

type Service interface {
	Greet(name string) string
}

type Repository interface {
	FindByName(name string) User
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Greet(name string) string {
	var user User

	if name != "" {
		user = s.repo.FindByName(name)
	}

	return sayHello(user)
}

func sayHello(user User) string {
	if user.Name == "" {
		return "Hello, world!"
	} else {
		return "Hello, " + user.Name + "!"
	}
}
