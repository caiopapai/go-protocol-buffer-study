package main

import (
	"fmt"

	simple "github.com/caiopapai/go-protocol-buffer-study/src"
)

func main() {
	doSimple()
}

func doSimple() {
	sa := simple.SimpleMessage{
		Id:         12345,
		Name:       "Caio Papai",
		IsSimple:   true,
		SimpleList: []int32{4, 5, 6, 7},
	}

	fmt.Println("Mensagem: ", sa)
}
