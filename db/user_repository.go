package repository

import "github.com/Arturlima/crud-go/models"

type IUserRepository interface {
	FindAll() ([]models.User, error)
	FindById(Id string) (models.User, error)
	Insert(user *models.User) (models.User, error)
	Update(Id string) (models.User, error)
	Delete(id string) error
}

type UserRepository struct {
	context *map[string]models.User
}

func NewUserRepository(u *map[string]models.User) IUserRepository {
	return &UserRepository{
		context: u,
	}
}

// Delete implements IUserRepository.
func (u *UserRepository) Delete(id string) error {
	panic("unimplemented")
}

// FindAll implements IUserRepository.
func (u *UserRepository) FindAll() ([]models.User, error) {
	panic("unimplemented")
}

// FindById implements IUserRepository.
func (u *UserRepository) FindById(Id string) (models.User, error) {
	panic("unimplemented")
}

// Insert implements IUserRepository.
func (u *UserRepository) Insert(user *models.User) (models.User, error) {
	panic("unimplemented")
}

// Update implements IUserRepository.
func (u *UserRepository) Update(Id string) (models.User, error) {
	panic("unimplemented")
}
