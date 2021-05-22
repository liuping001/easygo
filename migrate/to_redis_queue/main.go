// Author: coolliu
// Date: 2021/5/21

package main

import (
	"flag"
	"fmt"
	"github.com/liuping001/easygo/migrate/common"
	. "github.com/liuping001/fastgo/app"
	myredis "github.com/liuping001/fastgo/util/redis"
	"sync/atomic"
	"time"
)

var (
	queueKey = flag.String("queueKey", "", "redis 队列")
	queueNum = flag.Uint64("queueNum", 1, "redis 队列个数")
)

type toRedisQueue struct {
	i uint64
}

func (t *toRedisQueue) ConsumerMsg(msg interface{}) {
	value := string(msg.([]byte))

	index := atomic.AddUint64(&t.i, 1) % *queueNum
	key := fmt.Sprintf("%s:%d", *queueKey, index)

	ret := myredis.Client.LPush(key, value)
	if ret.Err() != nil {
		common.MyLog.Errorf("lpush error:%s\n", ret.Err())
		time.Sleep(1 * time.Second)
	}
}

type MySignal struct {
	OnSignal
	work *common.Worker
}

func (s *MySignal) OnExit() {
	s.work.OnExit()
}

func main() {
	flag.Parse()
	myredis.Init()

	work := common.Worker{
		Consumer: &toRedisQueue{},
	}

	GracefulExit(&MySignal{&DefaultOnSignal{}, &work}, common.MyLog)
	work.Run()
}
