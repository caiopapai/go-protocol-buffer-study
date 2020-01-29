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
	readFromFile("simple.bin", sm)
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatal("Error ", err.Error())
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatal("Erro ", err.Error())
		return err
	}

	fmt.Println("Written")
	return nil
}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatal("Error ", err.Error())
		return err
	}

	err2 := proto.Unmarshal(in, pb)
	if err != nil {
		log.Fatal("Error ", err.Error())
		return err2
	}

	fmt.Println("Message From File", pb)

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
