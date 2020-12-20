package repository

import "week04/internal/domain"

type Repository interface {
	FindBy(id int64) (*domain.User, error)
}
