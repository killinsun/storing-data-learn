package main

import (
	"fmt"
	"io/ioutil"
	"os"
)



func main() {

	data := []byte("Hello World!\n")
	err := ioutil.WriteFile("data1", data, 0644) // 0644 はパーミッション
	if err != nil {
		panic(err)
	}

	read1, _ := ioutil.ReadFile("data1")
	fmt.Print(string(read1))

	// ------------------------------------------
	// 構造体 File を使ったパターン
	// ------------------------------------------

	file1, _ := os.Create("data2") // 構造体Fileを取得
	defer file1.Close() //

	data2 := []byte("Hey! What's up?")

	bytes, _ := file1.Write(data2)
	fmt.Printf("Wrote %d bytes to file\n", bytes)

	file2, _ := os.Open("data2")
	defer file2.Close()

	read2 := make([]byte, len(data2))
	bytes, _ = file2.Read(read2)
	fmt.Printf("Read %d bytes from file\n", bytes)
	fmt.Println(string(read2))

}