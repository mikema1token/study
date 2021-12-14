package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

func NewDemoListener() *DemoListener {
	return &DemoListener{
		localtrans: new(sync.Map),
	}
}

type DemoListener struct {
	localtrans       *sync.Map
	transactionIndex int32
}

func (d *DemoListener) ExecuteLocalTransaction(message *primitive.Message) primitive.LocalTransactionState {
	index := atomic.AddInt32(&d.transactionIndex, 1)
	d.localtrans.Store(message.TransactionId, primitive.LocalTransactionState(1+index%3))
	return primitive.UnknowState
}

func (d *DemoListener) CheckLocalTransaction(ext *primitive.MessageExt) primitive.LocalTransactionState {
	v, ok := d.localtrans.Load(ext.TransactionId)
	if !ok {
		return primitive.CommitMessageState
	}
	state := v.(primitive.LocalTransactionState)
	switch state {
	case 1:
		return primitive.CommitMessageState
	case 2:
		return primitive.RollbackMessageState
	case 3:
		return primitive.UnknowState
	default:
		return primitive.CommitMessageState
	}
}

func main() {
	p, _ := rocketmq.NewTransactionProducer(NewDemoListener(),
		producer.WithNameServer([]string{"127.0.0.1:9876"}),
		producer.WithRetry(1))
	p.Start()
	for i := 0; i < 10; i++ {
		res, err := p.SendMessageInTransaction(context.Background(), primitive.NewMessage("test2", []byte("hello,rocketmq again"+strconv.Itoa(i))))
		if err != nil {
			fmt.Printf("send message error: %s\n", err)
		} else {
			fmt.Printf("send message success: result=%s\n", res.String())
		}
	}
	time.Sleep(5 * time.Minute)
	err := p.Shutdown()
	if err != nil {
		fmt.Printf("shutdown producer error: %s", err.Error())
	}
}
