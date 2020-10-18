package main

import "os"

func main() {
	text := []byte("default")

	file, err := os.Create(`../Go_gin_Chat/test.txt`)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.Write(text)
}
