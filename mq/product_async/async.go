package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"os"
	"strconv"
	"sync"
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
	group := sync.WaitGroup{}
	for i := 1000; i < 2000; i++ {
		group.Add(1)
		err = p.SendAsync(context.Background(), func(ctx context.Context, result *primitive.SendResult, err error) {
			if err == nil {
				group.Done()
				fmt.Printf("send message success,%s", result.String())
			} else {
				fmt.Printf("send message fail,%s", err.Error())
			}
		}, &primitive.Message{
			Topic: topic,
			Body:  []byte("hello,rocket go client!" + strconv.Itoa(i)),
		})
		if err != nil {
			return
		}
	}
	group.Wait()
	err = p.Shutdown()
	if err != nil {
		fmt.Printf("shutdown err:%s\n", err.Error())
	}

}
