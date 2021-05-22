// Author: coolliu
// Date: 2021/5/22

package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/liuping001/easygo/draw"
	"net/http"
	"strings"
	"time"
)

func show(w http.ResponseWriter, r *http.Request) {
	cmd := r.URL.Query().Get("cmd")
	if cmd == "" {
		fmt.Fprint(w, "cmd empty. use like ?cmd=\"key *\"\n")
		return
	}

	cmds := strings.Fields(cmd)

	if len(cmds) <= 0 {
		fmt.Fprint(w, "cmd wrong:"+cmd)
		return
	}

	if !checkCmd(cmds[0]) {
		fmt.Fprint(w, "cmd not allow:"+cmd)
		return
	}

	var args []interface{}
	for _, value := range cmds {
		args = append(args, value)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var ret *redis.Cmd
	switch redisClient.(type) {
	case *redis.Client:
		ret = redisClient.(*redis.Client).Do(ctx, args...)
	case *redis.ClusterClient:
		ret = redisClient.(*redis.ClusterClient).Do(ctx, args...)
	}
	if ret.Err() != nil && ret.Err() != redis.Nil {
		fmt.Fprint(w, fmt.Sprintf("redis do cmd error:{%s} cmds:%v\n", ret.Err().Error(), cmds))
		return
	}

	val := ret.Val()
	data, err := draw.DrawSvg(&val)
	if err != nil {
		fmt.Fprint(w, fmt.Sprintf("val:%v, err:%s", val, err.Error()))
		return
	}
	fmt.Fprint(w, data)
}

var allowCmd = map[string]bool{
	"scan":          true,
	"type":          true,
	"ttl":           true,
	"get":           true,
	"mget":          true,
	"lrange":        true,
	"smembers":      true,
	"hscan":         true,
	"sscan":         true,
	"hget":          true,
	"hmget":         true,
	"hgetall":       true,
	"zrange":        true,
	"zrevrange":     true,
	"zrank":         true,
	"zrevrank":      true,
	"zscore":        true,
	"zcard":         true,
	"zrangebyscore": true,
	"xrange":        true,
	"keys":          true,
	"info":          true,
}

func checkCmd(cmd string) bool {
	_, ok := allowCmd[cmd]
	return ok
}
