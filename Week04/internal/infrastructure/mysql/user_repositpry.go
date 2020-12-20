package mysql

import (
	"github.com/go-xorm/xorm"
	"log"
	"week04/internal/domain"
)

type UserRepository struct {
	engine *xorm.Engine
}

func (u *UserRepository) FindBy(id int64) (*domain.User, error) {
	user := domain.User{}
	err := u.engine.Where("name = ?", id).Find(&user)
	if err != err {
		log.Println(err)
	}
	return &user, err
}

func NewUserRepo(engine *xorm.Engine) *UserRepository {
	return &UserRepository{
		engine: engine,
	}
}
