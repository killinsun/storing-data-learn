package main

import (
	"fmt"
)

type Post struct {
	Id int
	Content string
	Author string
}

// map = 連想配列、ハッシュ
var PostById map[int] *Post // => {1: 0xc000098210, 2: 0xc...., 3: 0x....} 
var PostsByAuthor map[string] []*Post // {"Sau Sheong": 0xc000098210, ,,,}

func store(post Post) {
	PostById[post.Id] = &post
	PostsByAuthor[post.Author] = append(PostsByAuthor[post.Author], &post)
}

func main() {

	// サイズの指定を省略しているので、最小の開始サイズを割り当てたスライスを返す
	PostById = make(map[int] *Post)
	PostsByAuthor = make(map[string][] *Post)

	post1 := Post{Id: 1, Content: "Hello World!", Author: "Sau Sheong"}
	fmt.Println(PostById)
	fmt.Println(PostsByAuthor)
	post2 := Post{Id: 2, Content: "Bonjour Monde!", Author: "Pierre"}
	fmt.Println(PostById)
	fmt.Println(PostsByAuthor)
	post3 := Post{Id: 3, Content: "Hola Mundo!", Author:  "Pedro"}
	post4 := Post{Id: 4, Content: "Greetings Earthlings!", Author: "Sau Sheong"}

	fmt.Println(PostById[0])
	fmt.Println(PostById)

	store(post1)
	store(post2)
	store(post3)
	store(post4)

	fmt.Println(PostById)
	fmt.Println(PostsByAuthor)
	fmt.Println(PostById[1])
	fmt.Println(PostById[2])

	for _, post := range PostsByAuthor["Sau Sheong"] {
		fmt.Println(post)
	}

	for _, post := range PostsByAuthor["Pedro"] {
		fmt.Println(post)
	}
}