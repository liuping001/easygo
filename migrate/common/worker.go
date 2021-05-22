// Author: coolliu
// Date: 2021/5/21

package common

import (
	"context"
	"flag"
	"sync"
)

var (
	queueLen  = flag.Int("msgQueueLen", 5000, "生产者队列长度")
	workerNum = flag.Int("workerNum", 1, "工作协程数")
)

type Consumer interface {
	ConsumerMsg(msg interface{})
}

type Worker struct {
	Consumer
	group  sync.WaitGroup
	cancel context.CancelFunc
	ctx    context.Context
}

func (w *Worker) work(msgQ <-chan interface{}) {
	for msg := range msgQ {
		w.ConsumerMsg(msg)
	}
	w.group.Done()
}

func (w *Worker) Run() {
	w.ctx, w.cancel = context.WithCancel(context.Background())
	msgQueue := make(chan interface{}, *queueLen)
	producer := StdioProducer{
		Msg: msgQueue,
	}
	producer.BeginWork(w.ctx)
	for i := 0; i < *workerNum; i++ {
		w.group.Add(1)
		w.work(producer.MsgChan())
	}
	w.group.Wait()
}

func (w *Worker) OnExit() {
	w.cancel()
	MyLog.Infof("work will exit\n")
}
