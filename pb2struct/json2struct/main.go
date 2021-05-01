// Author: coolliu
// Date: 2021/4/24

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	file       = flag.String("file", "", "json")
	structType = flag.String("to_type", "rapidjson", "watch cpp json lib. rapidjson")
)

func ReadAll(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	return ioutil.ReadAll(file)
}

func toRipadJson(data []byte) {
	{
		t := JsonOutStruct{StructOutI(&CppStructOut{}),
			TypeTransformI(&CppTypeTransform{}),
			[]string{}}
		t.Object([]byte("root"), []byte(data))
	}
	{
		t := JsonOutParseFunc{OutJsonFuncI(&RapidJsonParseFuncOut{}),
			TypeTransformI(&CppTypeTransform{}),
			[]string{}}
		t.Object([]byte("root"), []byte(data))
	}
}

func main() {
	flag.Parse()

	json, err := ReadAll(*file)
	if err != nil {
		fmt.Print("err:%s", err.Error())
		os.Exit(1)
		return
	}
	if *structType == "rapidjson" {
		toRipadJson(json)
	}

}
