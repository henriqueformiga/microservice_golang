package domain

type User struct {
	Name  string
	Email string
}

type UserRepository interface {
	FindAll() ([]User, error)
}

type UserRepositoryStub struct {
	users []User
}

func (s UserRepositoryStub) FindAll() ([]User, error) {
	return s.users, nil
}
