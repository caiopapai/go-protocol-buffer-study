package main

import (
	"fmt"
	"io/ioutil"
	"log"

	simple "github.com/caiopapai/go-protocol-buffer-study/src"
	"github.com/golang/protobuf/proto"
)

func main() {
	sm := doSimple()
	writeToFile("simple.bin", sm)
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatal("Error ", err)
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatal("Erro ", err)
		return err
	}

	fmt.Println("Written")
	return nil
}

func doSimple() *simple.SimpleMessage {
	sm := simple.SimpleMessage{
		Id:         12345,
		Name:       "Caio Papai",
		IsSimple:   true,
		SimpleList: []int32{4, 5, 6, 7},
	}

	fmt.Println("Mensagem: ", sm)

	return &sm
}
