package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Post struct {
	Id int
	Content string
	Author string `sql:"not null"`
	Comments []Comment
	CreatedAt time.Time
}

type Comment struct {
	Id int
	Content string
	Author string `sql:"not null"`
	PostId int
	CreatedAt time.Time
}

var Db *gorm.DB

func init() {
	var err error
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset-utf8&parseTime=true", "test","test", "127.0.0.1", "gwp")
	Db, err = gorm.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&Post{}, &Comment{})
}

func main() {
	post := Post{Content: "Hello World!", Author: "Sau Sheong"}
	fmt.Println(post)

	Db.Create(&post)
	fmt.Println(post)

	comment := Comment{Content: "良い投稿だね！", Author: "Joe"}
	Db.Model(&post).Association("Comments").Append(comment)

	var readPost Post
	Db.Where("author = ?", "Sau Sheong").First(&readPost)
	var comments []Comment
	Db.Model(&readPost).Related(&comments)
	fmt.Println(comments[0])

}