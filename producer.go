package main

import (
	"encoding/json"
	"log"

	"github.com/nsqio/go-nsq"
)

type Message struct {
	Name      string
	Content   string
	Timestamp string
}

func main() {
	//The only valid way to instantiate the Config
	config := nsq.NewConfig()
	//Creating the Producer using NSQD Address
	producer, err := nsq.NewProducer("172.18.13.123:4150", config)
	if err != nil {
		log.Fatal(err)
	}

	//Init topic name and message
	topic := "Topic"
	msg := Message{
		Name:      "Message",
		Content:   "Hello World",
		Timestamp: "2021-09-29",
	}
	//Convert message as []byte
	payload, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
	}
	//Publish the Message
	err = producer.Publish(topic, payload)
	if err != nil {
		log.Println(err)
	}
}
