package domain

type User struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
}

type UserRepository interface {
	GetByID(id int) (*User, error)
	Save(user *User) error
}
