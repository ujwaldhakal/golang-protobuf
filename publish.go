package main

import (
	"encoding/json"
	"fmt"
	proto "github.com/ujwaldhakal/golang-protobuf/messages/proto"
	rabbitmq "github.com/ujwaldhakal/golang-protobuf/pkg"
	"log"
)

func main()  {

	data := proto.UserRegistered{
		UserId: 1,
		Email:  "johndoe@example.com",
	}

	encodedData, err := json.Marshal(data)
	if err != nil {
		log.Fatal("error encoding to json")
	}
	rabbitmq.Publish("UserRegistered", encodedData)
	fmt.Println("Data has been published")
}