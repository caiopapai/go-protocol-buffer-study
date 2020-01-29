package main

import (
	"fmt"
	"io/ioutil"
	"log"

	enumspb "github.com/caiopapai/go-protocol-buffer-study/enum"
	simple "github.com/caiopapai/go-protocol-buffer-study/src"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func main() {
	sm := doSimple()
	sm2 := &simple.SimpleMessage{}

	writeToFile("simple.bin", sm)
	readFromFile("simple.bin", sm2)

	smString := toJSON(sm2)
	fmt.Println("Protocol Buffer to JSON: ", smString)

	sm3 := &simple.SimpleMessage{}
	fromJSON(smString, sm3)
	fmt.Println("JSON to Protocol Buffer: ", sm3)

	doEnum()

}

func toJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatal("Error ", err.Error())
	}
	return out
}

func fromJSON(json string, pb proto.Message) {
	err := jsonpb.UnmarshalString(json, pb)
	if err != nil {
		log.Fatal("Error ", err.Error())
	}

}

func doEnum() {
	en := enumspb.EnumMessage{
		Id:        32,
		DayOfWeek: enumspb.DayOfWeek_QUARTA,
	}

	fmt.Println(en)
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

	fmt.Println("Write To File :", out)
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

	fmt.Println("Message From File: ", pb)

	return nil
}

func doSimple() *simple.SimpleMessage {
	sm := simple.SimpleMessage{
		Id:         12345,
		Name:       "Caio Papai",
		IsSimple:   true,
		SimpleList: []int32{4, 5, 6, 7},
	}

	fmt.Println("Protocol Buffer Message: ", sm)

	return &sm
}
