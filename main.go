package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {

	h := sha256.New()

	h.Write([]byte("Test message for hash"))
	fmt.Printf("%x", h.Sum(nil))

}
