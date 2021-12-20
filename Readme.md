## Use Protobuf with Go

This example demonstrates on how to use [Protobuf](https://developers.google.com/protocol-buffers) with Go.


### Installation 
* Clone this repo
* Run `make start` to initialize the docker containers
* Run `make build-proto` to convert the proto files into auto generated code.
* Run `make publish` to publish queue on rabbitmq
* Run `make consume` to consume queue

### Technologies Used
* [Protobuf](https://developers.google.com/protocol-buffers)
* [GO](https://go.dev/)
* [Rabitmq](https://www.rabbitmq.com/)