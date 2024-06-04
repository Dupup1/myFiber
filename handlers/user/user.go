package user

import (
	"fmt"
	"github.com/Dupup1/myFiber/database"
	"github.com/Dupup1/myFiber/models"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	db := database.DBConn
	var users []models.User
	db.Find(&users)
	return c.JSON(fiber.Map{
		"success": true,
		"users":   users,
	})
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var user models.User
	db.Find(&user, id)
	if user.Name == "" {
		return c.Status(500).SendString("没找到该用户")
	}
	return c.JSON(fiber.Map{
		"success": true,
		"users":   user,
	})

}

func NewUser(c *fiber.Ctx) error {
	db := database.DBConn
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).SendString("用户加载失败")
	}
	db.Create(&user)
	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var user models.User
	db.First(&user, id)
	if user.Name == "" {
		return c.Status(500).SendString("没找到该用户")
	}
	db.Delete(&user)
	fmt.Println("删除成功")
	return c.JSON(user)
}
