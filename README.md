# myFiber

> 技术栈 golang, Go-Fiber, Gorm, mysql

## 使用
1. 导入mysql：`database/myfiber.sql`

2. `database/database.go`中修改`dsn`，连接至mysql
   ```go
   dsn := "username:password@tcp(127.0.0.1:3306)/myfiber?charset=utf8mb4&parseTime=True&loc=Local"
   ```

## 接口
| METHOD | ROUTE         | 作用         |
| ------ | ------------- | -------------- |
| GET    | /api/user     | 获取user表  |
| GET    | /api/user/:id | 获取id对应user |
| POST   | /api/user     | 增加用户   |
| DELETE | /api/user/:id | 删除id对应user |