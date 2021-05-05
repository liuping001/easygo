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
	outDir     = flag.String("out", "./", "out dir")
	pbType     = flag.String("pb_type", "json", "all pb type: json、yml、toml")
	structType = flag.String("to_type", "rapidjson", "watch cpp json lib. rapidjson")
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
			cpp_out.JsonToRapidJson(data, util.OutInfo{Path: *file, OutDir: *outDir})
		}
	}
}
