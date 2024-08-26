package repository

import (
	"errors"

	"github.com/Arturlima/crud-go/models"
)

type IUserRepository interface {
	FindAll() ([]models.User, error)
	FindById(Id string) (*models.User, error)
	Insert(user *models.User) (*models.User, error)
	Update(Id string) (models.User, error)
	Delete(id string) error
}

type UserRepository struct {
	context map[string]models.User
}

func NewUserRepository(u map[string]models.User) IUserRepository {
	return &UserRepository{
		context: u,
	}
}

func (u *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User

	for _, u := range u.context {
		users = append(users, u)
	}
	return users, nil
}

func (u *UserRepository) FindById(Id string) (*models.User, error) {
	user, err := u.context[Id]
	if !err {
		return nil, nil
	}
	return &user, nil
}

func (u *UserRepository) Insert(user *models.User) (*models.User, error) {
	u.context[user.ID] = *user
	if _, ok := u.context[user.ID]; !ok {
		return nil, errors.New("falied to insert user")
	}

	return user, nil
}

func (u *UserRepository) Update(Id string) (models.User, error) {
	panic("unimplemented")
}

func (u *UserRepository) Delete(id string) error {
	delete(u.context, id)
	return nil
}
