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

	file := "file.ext"

	dat, err := os.ReadFile(file)
	check(err)
	fmt.Print(string(dat))

}
