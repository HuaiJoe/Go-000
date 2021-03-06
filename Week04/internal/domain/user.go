package domain

import "time"

type User struct {
	Id      int64
	Name    string
	Passwd  string    `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}
