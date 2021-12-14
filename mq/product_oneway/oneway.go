package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"os"
	"strconv"
)

func main() {
	p, _ := rocketmq.NewProducer(
		producer.WithNameServer([]string{"127.0.0.1:9876"}),
		producer.WithRetry(2),
		producer.WithGroupName("producer_test"),
	)
	err := p.Start()
	if err != nil {
		fmt.Printf("start producer error:%s", err.Error())
		os.Exit(1)
	}
	topic := "test"
	for i := 2000; i < 3000; i++ {
		err = p.SendOneWay(context.Background(), &primitive.Message{
			Topic: topic,
			Body:  []byte("hello,rocket go client!" + strconv.Itoa(i)),
		})
		if err != nil {
			fmt.Printf("send message oneway,%s", err.Error())
		}
	}
	err = p.Shutdown()
	if err != nil {
		fmt.Printf("shutdown err:%s\n", err.Error())
	}
}
