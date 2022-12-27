package consumer

import (
	"context"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"mygolangapp/internal/infra/kafka/factory"
	"mygolangapp/pkg/uow"
)

func Consume(topics []string, servers string, msgChan chan *kafka.Message) {
    kafkaConsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
        "bootstrap.servers": servers,
        "group.id":          "gostats",
        "auto.offset.reset": "earliest",
    })
    if err != nil {
        panic(err)
    }
    kafkaConsumer.SubscribeTopics(topics, nil)
    for {
        msg, err := kafkaConsumer.ReadMessage(-1)
        if err == nil {
            msgChan <- msg
        }
    }
}

func ProcessEvents(ctx context.Context, msgChan chan *kafka.Message, uwo uow.UowInterface) {
    for msg := range msgChan {
        fmt.Println("Received message", string(msg.Value), "on topic", *msg.TopicPartition.Topic)
        strategy := factory.CreateProcessMessageStrategy(*msg.TopicPartition.Topic)
        err := strategy.Process(ctx, msg, uwo)
        if err != nil {
            fmt.Println(err)
        }
    }
}