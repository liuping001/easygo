// Author: coolliu
// Date: 2021/5/22

package main

import (
	"flag"
	"github.com/go-redis/redis/v7"
	"github.com/liuping001/easygo/migrate/common"
	. "github.com/liuping001/fastgo/app"
	myredis "github.com/liuping001/fastgo/util/redis"
	"strings"
	"time"
)

type toRedis struct {
	i uint64
}

func (t *toRedis) ConsumerMsg(msg interface{}) {
	value := string(msg.([]byte))
	if value == "" {
		common.MyLog.Infof("msg empty\n")
		return
	}
	cmds := strings.Fields(value)
	var args []interface{}
	for _, value := range cmds {
		args = append(args, value)
	}
	var ret *redis.Cmd
	switch myredis.Client.(type) {
	case *redis.Client:
		ret = myredis.Client.(*redis.Client).Do(args...)
	case *redis.ClusterClient:
		ret = myredis.Client.(*redis.ClusterClient).Do(args...)
	}
	if ret.Err() != nil {
		common.MyLog.Infof("redis do cmd error:{%s} cmds:%v\n", ret.Err().Error(), cmds)
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
		Consumer: &toRedis{},
	}

	GracefulExit(&MySignal{&DefaultOnSignal{}, &work}, common.MyLog)
	work.Run()
}
