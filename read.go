package main

import (
	// "bufio"

	"fmt"
	// "io"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
}

func read() {

	file := "assets/all.json"

	dat, err := os.ReadFile(file)

	check(err)

	fmt.Print(string(dat))

	// var albums = []Album
	// error := json.Unmarshal(dat, albums)
	// check(error)
	// if error != nil {
	// 	fmt.Println("error:", error)
	// }

	// fmt.Printf("%+v", albums)
}
