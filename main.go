package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Post struct {
	Id int
	Content string
	AuthorName string `db: autho`
}

var Db *sqlx.DB

func init() {
	var err error
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset-utf8", "test","test", "127.0.0.1", "gwp")
	Db, err = sqlx.Open("mysql", dataSourceName)
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
		err = rows.Scan(&post.Id, &post.Content, &post.AuthorName)
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
	err = Db.QueryRowx("select id, content, author from posts where id = ?", id).StructScan(&post)
	if err != nil {
		return
	}
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
	result, err := stmt.Exec(post.Content, post.AuthorName)
	lid, err := result.LastInsertId()

	post.Id = int(lid)
	return 
}

func (post *Post) Update() (err error) {
	_, err = Db.Exec("update posts set content = ?, author = ? where id = ?", post.Content, post.AuthorName, post.Id)
	return
}

func (post *Post) Delete() (err error) {
	_, err = Db.Exec("delete from posts where id = ?", post.Id)
	return
}


func main() {
	post := Post{Content: "Hello World!", AuthorName : "Sau Sheong"}

	post.Create()

	fmt.Println(post)

	readPost := Post{}
	readPost, _ = GetPost(post.Id)
	fmt.Println(readPost)

}