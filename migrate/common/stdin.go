// Author: coolliu
// Date: 2021/5/21

package common

import (
	"bufio"
	"context"
	"os"
)

// 从标准输入读取一行行数据
type Producer interface {
	Msg() <-chan interface{}
}

type StdioProducer struct {
	Msg chan interface{}
}

func (s *StdioProducer) MsgChan() <-chan interface{} {
	return s.Msg
}
func (s *StdioProducer) BeginWork(ctx context.Context) {
	go s.readAllLine(ctx)
}

func (s *StdioProducer) readAllLine(ctx context.Context) {
	defer close(s.Msg)
	input := bufio.NewReader(os.Stdin)
	for {
		data, _, err := input.ReadLine()
		if err != nil {
			return
		}
		s.Msg <- data
		select {
		case <-ctx.Done():
			return
		default:
		}
	}
}
