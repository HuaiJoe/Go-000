package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"time"
)

type User struct {
	Id      int64
	Name    string
	Passwd  string    `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

// 1 业务层在查询DB的时候返回了错误，建议封装一下，底层的错误信息不应该直接宝楼给用户
// 2 如果是中间件层对中间点层或者第三方库的调用，建议直接return 原始err
func (u *User) Find(engine *xorm.Engine) error {
	err := engine.Where("name = ?", u.Name).Find(u)
	if err != nil {
		return fmt.Errorf("can not find the record from db,%w", err)
	}
	return nil
}

func CreateDBEngine(ip, port, user, passwd string) (*xorm.Engine, error) {
	cmd := fmt.Sprintf("%v:%v@tcp(%v:%v)/test?charset=utf8mb4", user, passwd, ip, port)
	return xorm.NewEngine("mysql", cmd)
}

func main() {
	engine, err := CreateDBEngine("localhost", "3306", "root", "Happy100")

	if err != nil {
		fmt.Print("error")
		log.Fatal("create connections failed.", err)
	}

	defer engine.Close()

	user := User{}
	err = user.Find(engine)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(user)
}
