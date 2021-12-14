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
		consumer.WithGroupName("consumer_flag"),
		consumer.WithConsumeFromWhere(consumer.ConsumeFromLastOffset),
	)
	if err != nil {
		log.Fatal(err)
	}
	selector := consumer.MessageSelector{
		Type:       consumer.TAG,
		Expression: "taga||tagb",
	}
	c.Subscribe("flag", selector, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		totalC := 0
		totalAB := 0
		for i := 0; i < len(msgs); i++ {
			if msgs[i].GetTags() == "tagc" {
				totalC++
			} else {
				totalAB++
			}
		}
		fmt.Printf("subscibe callback:%v,totalC:%d,totalAB:%d \n", msgs, totalC, totalAB)
		return consumer.ConsumeSuccess, nil
	})
	c.Start()
	time.Sleep(time.Hour)
	c.Shutdown()
}
