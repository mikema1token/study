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
		consumer.WithGroupName("order_test_2"),
		consumer.WithConsumerOrder(true),
		consumer.WithConsumeFromWhere(consumer.ConsumeFromFirstOffset),
	)
	if err != nil {
		log.Fatal(err)
	}
	c.Subscribe("test_order_3", consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		orderlyCtx, b := primitive.GetOrderlyCtx(ctx)
		orderlyCtx.AutoCommit = false
		fmt.Printf("orderlyCtx:%v,exist:%v,length:%d", orderlyCtx, b, len(msgs))
		fmt.Printf("subscibe callback:%v \n", msgs)
		for _, v := range msgs {
			fmt.Println(v.GetKeys())
		}

		return consumer.ConsumeSuccess, nil
	})
	c.Start()
	time.Sleep(time.Hour)
	c.Shutdown()
}
