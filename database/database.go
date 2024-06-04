package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func InitDatabase() error {
	dsn := "root:root@tcp(127.0.0.1:3306)/myfiber?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("数据库打开失败")
	}
	fmt.Println("数据库打开成功")
	return nil
}
