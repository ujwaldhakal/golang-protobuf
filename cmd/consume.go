package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	v1 "github.com/ujwaldhakal/golang-protobuf/message/proto/v1"
	"github.com/ujwaldhakal/golang-protobuf/pkg/rabbitmq"
	"log"
)

func init() {
	rootCmd.AddCommand(consume)
}

var consume = &cobra.Command{
	Use:   "consume",
	Short: "Try and possibly fail at something",
	Run: func(cmd *cobra.Command, args []string) {

		msgs := rabbitmq.ConsumerClient()

		forever := make(chan bool)

		go func() {
			for d := range msgs {

				jsonMap := &v1.UserRegistered{}

				 err := json.Unmarshal(d.Body, jsonMap)
				 if err != nil {
					 fmt.Println("error decoding json")
				 }

				 errors := validateMessage(jsonMap)

				 if errors != nil {
					 log.Fatalf("UserRegistered-validation-error: ",errors)
				 }

				fmt.Println(jsonMap)
			}
		}()

		log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
		<-forever
	},

}

func validateMessage(eventPayload *v1.UserRegistered) error {
	p := new(v1.UserRegistered)
	p.Email = eventPayload.Email
	p.UserId = eventPayload.UserId

	return p.ValidateAll()
}
