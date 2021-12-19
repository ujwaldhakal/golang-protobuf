package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	proto "github.com/ujwaldhakal/golang-protobuf/message/proto/v1"
	"github.com/ujwaldhakal/golang-protobuf/pkg/rabbitmq"
	"log"
)

func init() {
	rootCmd.AddCommand(publish)
}

var publish = &cobra.Command{
	Use:   "publish",
	Short: "Try and possibly fail at something",
	Run: func(cmd *cobra.Command, args []string) {

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
	},
}
