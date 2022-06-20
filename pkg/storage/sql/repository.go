package sql

import (
	"database/sql"
	"github.com/maccath/go-webapp/pkg/greeting"
)

type Repository interface {
	FindByName(name string) greeting.User
	Save(user User)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

func (r repository) FindByName(name string) greeting.User {
	var user greeting.User
	var id string
	var n string

	row := r.db.QueryRow("SELECT id, name FROM users WHERE name = '" + name + "'")

	row.Scan(&id, &n)

	if n == name {
		user.Name = n
	}

	return user
}

func (r repository) Save(user User) {
	// Todo
}
