package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Id int
	Content string
	Author string
}

var Db *sql.DB

func init() {
	var err error
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset-utf8", "test","test", "127.0.0.1", "gwp")
	Db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
}

func Posts(limit int) (posts []Post, err error) {
	rows, err := Db.Query("select id, content, author from posts limit ?", limit)
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("select id, content, author from posts where id = ?", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

func (post *Post) Create() (err error) {
	statement := "insert into posts (content, author) values (?, ?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer stmt.Close()
	result, err := stmt.Exec(post.Content, post.Author)
	lid, err := result.LastInsertId()

	post.Id = int(lid)
	return 
}

func (post *Post) Update() (err error) {
	_, err = Db.Exec("update posts set content = ?, author = ? where id = ?", post.Content, post.Author, post.Id)
	return
}

func (post *Post) Delete() (err error) {
	_, err = Db.Exec("delete from posts where id = ?", post.Id)
	return
}


func main() {
	post := Post{Content: "Hello World!", Author : "Sau Sheong"}

	post.Create()

	readPost, _ := GetPost(post.Id)
	fmt.Println(readPost)

	readPost.Content = "Bonjour Monde!"
	readPost.Author = "Pierre"
	readPost.Update()

	posts, _ := Posts(10)
	fmt.Println(posts)

	readPost.Delete()

}