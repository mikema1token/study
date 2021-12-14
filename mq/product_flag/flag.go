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
		producer.WithGroupName("product_flag"),
	)
	err := p.Start()
	if err != nil {
		fmt.Printf("start producer error:%s", err.Error())
		os.Exit(1)
	}
	topic := "flag"
	flags := []string{"taga", "tagb", "tagc"}
	for i := 0; i < 100; i++ {
		msg := primitive.Message{
			Topic: topic,
			Body:  []byte("hello,rocket go client!" + strconv.Itoa(i)),
		}
		msg.WithTag(flags[i%3])
		r, err := p.SendSync(context.Background(), &msg)
		if err != nil {
			fmt.Printf("send message fail,%s\n", err.Error())
		} else {
			fmt.Printf("send message success,%s\n", r.String())
		}
	}
	err = p.Shutdown()
	if err != nil {
		fmt.Printf("shutdown err:%s\n", err.Error())
	}
}
