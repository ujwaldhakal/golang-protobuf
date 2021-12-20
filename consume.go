package main

import (
	"encoding/json"
	"fmt"
	proto "github.com/ujwaldhakal/golang-protobuf/messages/proto"
	"github.com/ujwaldhakal/golang-protobuf/pkg"
	"log"
)

func main()  {

	msgs := pkg.ConsumerClient("UserRegistered")

	forever := make(chan bool)

	go func() {
		for d := range msgs {

			jsonMap := &proto.UserRegistered{}

			err := json.Unmarshal(d.Body, jsonMap)
			if err != nil {
				fmt.Println("error decoding json")
			}

			errors := validateMessage(jsonMap)

			if errors != nil {
				log.Fatalf("UserRegistered-validation-error: ", errors)
			}

			fmt.Println(jsonMap)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}



func validateMessage(eventPayload *proto.UserRegistered) error {
	p := new(proto.UserRegistered)
	p.Email = eventPayload.Email
	p.UserId = eventPayload.UserId

	return p.ValidateAll()
}
