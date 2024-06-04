package migration

import (
	"fmt"
	"github.com/Dupup1/myFiber/database"
	"github.com/Dupup1/myFiber/models"
	"log"
)

func RunMigration() {
	err := database.DBConn.AutoMigrate(&models.User{})
	if err != nil {
		log.Println(err)
		fmt.Println("数据库迁移失败")
	}
	fmt.Println("数据库迁移完成")
}
