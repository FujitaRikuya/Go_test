package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	if err := readBytes("read.txt"); err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}

}

func readBytes(filename string) error {
	file, err := os.Open(filename)

	if err != nil {
		return err
	}

	defer file.Close()

	b := make([]byte, 10)

	//おそらく無限ループ(whileみたいな使い方)
	for {
		c, err := file.Read(b)

		//error handling?????
		if c == 0 {
			break
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		line := string(b[:c])
		fmt.Print(line)
	}

	return nil
}
