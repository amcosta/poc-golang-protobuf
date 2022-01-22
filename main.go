package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/amcosta/poc-golang-protobuf/github.com/amcosta/poc-golang-protobuf/tutorialpb"
	"google.golang.org/protobuf/proto"
)

func main() {
	fmt.Println("poc golang protobuf")

	// person := pb.Person{
	// 	Id:    1234,
	// 	Name:  "John Doe",
	// 	Email: "jdoe@example.com",
	// 	Phones: []*pb.Person_PhoneNumber{
	// 		{Number: "555-4321", Type: pb.Person_HOME},
	// 	},
	// }

	// out, err := proto.Marshal(&person)
	// if err != nil {
	// 	log.Fatalln("Failed to encode message ", err)
	// }

	// if err := ioutil.WriteFile("pboutput.txt", out, 0644); err != nil {
	// 	log.Fatalln("Failed to write message in file ", err)
	// }

	in, err := ioutil.ReadFile("pboutput.txt")
	if err != nil {
		log.Fatalln(err)
	}

	var person tutorialpb.Person
	if err := proto.Unmarshal(in, &person); err != nil {
		log.Fatalln(err)
	}

	fmt.Println(&person)
}
