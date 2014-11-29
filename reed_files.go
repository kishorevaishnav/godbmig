package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Read files in the current directory")
	files, _ := ioutil.ReadDir("./")
	for _, f := range files {
		ReadFile(f.Name())
		fmt.Println(f.Name())
	}
}

func ReadFile(filename string) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)
}
