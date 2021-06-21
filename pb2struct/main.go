// Author: coolliu
// Date: 2021/4/24

package main

import (
	"flag"
	"fmt"
	"github.com/liuping001/easygo/pb2struct/cpp_out"
	"github.com/liuping001/easygo/pb2struct/util"
	"os"
)

var (
	file       = flag.String("file", "", "json")
	outDir     = flag.String("out_dir", "./", "out dir")
	pbType     = flag.String("pb_type", "json", "all pb type: json、yml、toml")
	structType = flag.String("to_type", "rapidjson", "cpp json lib. rapidjson")
	outFunc    = flag.Int("out_func", 0, "0:from,to 1:from 2:to")
)

func main() {
	flag.Parse()

	data, err := util.ReadAll(*file)
	if err != nil {
		fmt.Printf("err:%s", err.Error())
		os.Exit(1)
		return
	}
	if *pbType == "json" {
		if *structType == "rapidjson" {
			err := cpp_out.JsonToRapidJson(data, util.OutInfo{Path: *file, OutDir: *outDir, OutFunc: *outFunc})
			if err != nil {
				fmt.Errorf("%s", err.Error())
			}
		}
	} else {
		fmt.Errorf("没有定义的pbType：%s\n", *pbType)
		os.Exit(1)
	}
}
