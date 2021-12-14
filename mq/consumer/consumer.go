package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"log"
	"time"
)

func main() {
	c, err := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{"127.0.0.1:9876"}),
		consumer.WithConsumerModel(consumer.Clustering),
		consumer.WithGroupName("consumer_test"),
	)
	if err != nil {
		log.Fatal(err)
	}
	c.Subscribe("test2", consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		fmt.Printf("subscibe callback:%v \n", msgs)
		return consumer.ConsumeSuccess, nil
	})
	c.Start()
	time.Sleep(time.Hour)
	c.Shutdown()
}
