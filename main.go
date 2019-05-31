package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func main() {

	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer f.Close()

	h := sha256.New()
	io.Copy(h, f)

	fmt.Println("The sum is", h.Sum(nil))

}
