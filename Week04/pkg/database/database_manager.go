package database

import (
	"fmt"
	"github.com/go-xorm/xorm"
)

func NewDBEngine(ip, port, user, passwd string) (*xorm.Engine, error) {
	cmd := fmt.Sprintf("%v:%v@tcp(%v:%v)/test?charset=utf8mb4", user, passwd, ip, port)
	return xorm.NewEngine("mysql", cmd)
}
