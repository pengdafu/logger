package main

import "fmt"

type User struct {
	Id    uint
	Name  string
	Email string
	Age   uint8
}

func (user *User) TableName() string {
	return "user"
}

func (user *User) Insert() {
	//这里使用了Table()函数，如果你没有指定全局表名禁用复数，或者是表名跟结构体名不一样的时候
	//你可以自己在sql中指定表名。这里是示例，本例中这个函数可以去除。
	err := db.Create(user).Error
	fmt.Println("执行了insert", err)
}

func main() {
	user := &User{Name: "pdf", Age: 18, Email: "123@qq.com"}
	user.Insert()
}
