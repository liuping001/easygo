// Author: coolliu
// Date: 2021/3/21

package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/liuping001/easygo/draw"
	"net/http"
)

var (
	redisAddr = flag.String("redis_addr", "127.0.0.1:6379", "redis addr")
	port      = flag.String("port", "80", "http port")
)

var (
	redisClient *redis.Client
)

func main() {
	flag.Parse()

	// init redis
	redisClient = redis.NewClient(&redis.Options{
		Addr: *redisAddr,
	})

	http.HandleFunc("/", index)
	fmt.Printf("Server started at port %s\n", *port)
	http.ListenAndServe(fmt.Sprintf(":%s", *port), nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		fmt.Fprint(w, "key empty. use like ?key=mykey\n")
		return
	}
	kType := keyType(key)
	d := NewDrawByType(kType, key)
	if d == nil {
		fmt.Fprint(w, fmt.Sprintf("key:%s type:%s not support draw!\n", key, kType))
		return
	}
	data, err := d.draw()
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	fmt.Fprint(w, data)
}

func NewDrawByType(kType, key string) IDraw {
	switch kType {
	default:
		return nil
	case "stream":
		return &stream{key: key}
	}
}

func keyType(key string) string {
	ctx := context.Background()
	strRet := redisClient.Type(ctx, key)
	ret, err := strRet.Result()
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return ret
}

type Consumer struct {
	name    string
	pending int64
	idle    int64
}

type Group struct {
	name              string
	pending           int64
	last_delivered_id string
	consumers         []*Consumer
}

type Stream struct {
	length            int64
	redis_tree_keys   int64
	redis_tree_nodes  int64
	last_generated_id string
	groups            []*Group
}

type IDraw interface {
	draw() (string, error)
}

type stream struct {
	key string
}

func (s *stream) draw() (string, error) {
	var info Stream
	ctx := context.Background()

	streamRet := redisClient.XInfoStream(ctx, s.key)
	streamInfo, err := streamRet.Result()
	if err != nil {
		return "", err
	}
	info.length = streamInfo.Length
	info.redis_tree_keys = streamInfo.RadixTreeKeys
	info.redis_tree_nodes = streamInfo.RadixTreeNodes
	info.last_generated_id = streamInfo.LastGeneratedID

	groupsRet := redisClient.XInfoGroups(ctx, s.key)
	groupInfo, err := groupsRet.Result()
	if err != nil {
		return "", err
	}

	for _, item := range groupInfo {
		g := &Group{
			name:              item.Name,
			pending:           item.Pending,
			last_delivered_id: item.LastDeliveredID,
		}
		consumReq := redisClient.XInfoConsumers(ctx, s.key, item.Name)
		consumerInfo, err := consumReq.Result()
		if err != nil {
			return "", nil
		}
		for _, c := range consumerInfo {
			g.consumers = append(g.consumers, &Consumer{
				name:    c.Name,
				pending: c.Pending,
				idle:    c.Idle,
			})
		}
		info.groups = append(info.groups, g)
	}
	return draw.DrawSvg(&info)
}
