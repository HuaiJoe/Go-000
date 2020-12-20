package service

import (
	"week04/internal/domain"
	"week04/internal/repository"
	"week04/internal/view"
)

type UserCase interface {
	Register(user *domain.User) (int64, error)
	ChangePassword(user *domain.User) (int64, error)
	Query(id int64) (user view.UserView, err error)
}

type UserCaseService struct {
	repo *repository.Repository
}

func (u *UserCaseService) Register(user *domain.User) (int64, error) {
	//todo
	panic("implement me")
}

func (u *UserCaseService) ChangePassword(user *domain.User) (int64, error) {
	//todo
	panic("implement me")
}

func (u *UserCaseService) Query(id int64) (user view.UserView, err error) {
	data, err := u.Query(id)
	if err != nil {
		return
	}
	// todo refactor should use mapper function
	user.Id = data.Id
	user.Name = data.Name
	return
}

func NewUserCase(repository repository.Repository) *UserCaseService {
	return &UserCaseService{
		repo: &repository,
	}
}
