// Author: coolliu
// Date: 2021/4/24

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var (
	file       = flag.String("file", "", "json")
	outDir     = flag.String("out", "./", "out dir")
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

func FileName(file string) string {
	base := filepath.Base(file)
	suffix := path.Ext(base)
	return strings.TrimSuffix(base, suffix)
}

func NewFile(file string, dir string, suffix string) (*os.File, error) {
	fileName := fmt.Sprintf("%s/%s%s", dir, FileName(file), suffix)
	return os.OpenFile(
		fileName,
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
}

type OutInfo struct {
	path   string
	outDir string
}

func toRipadJson(data []byte, out OutInfo) {
	header := JsonOutStruct{StructOutI(&CppStructOut{}),
		TypeTransformI(&CppTypeTransform{}),
		[]string{}}
	header.Object([]byte("root"), data)

	body := JsonOutParseFunc{OutJsonFuncI(&RapidJsonParseFuncOut{}),
		TypeTransformI(&CppTypeTransform{}),
		[]string{}}
	body.Object([]byte("root"), data)

	h, err := NewFile(out.path, out.outDir, ".h")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer h.Close()
	h.WriteString(fmt.Sprintf("namespace %s{\n", FileName(out.path)))
	for _, item := range header.OutString {
		h.WriteString(item)
	}
	for _, item := range body.OutString {
		h.WriteString(item)
	}
	h.WriteString(fmt.Sprintf("\n}\n"))
}

func main() {
	flag.Parse()

	json, err := ReadAll(*file)
	if err != nil {
		fmt.Printf("err:%s", err.Error())
		os.Exit(1)
		return
	}
	if *structType == "rapidjson" {
		toRipadJson(json, OutInfo{*file, *outDir})
	}

}
