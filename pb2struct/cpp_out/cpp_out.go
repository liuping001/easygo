// Author: coolliu
// Date: 2021/5/5

package cpp_out

import (
	"fmt"
	"github.com/liuping001/easygo/pb2struct/common"
	"github.com/liuping001/easygo/pb2struct/json2struct"
	"github.com/liuping001/easygo/pb2struct/util"
)

type cppOutFuncI interface {
	common.OutFuncI
	Include() string
	RootFunc() string
}

func JsonToCpp(data []byte, out util.OutInfo, outFunc cppOutFuncI) {
	toStruct := json2struct.JsonOutStruct{
		OutStructI:     &CppStructOut{},
		TypeTransformI: &json2struct.CppTypeTransform{},
		OutString:      []string{},
	}
	toStruct.Object([]byte("root"), data)

	toFunc := json2struct.JsonOutParseFunc{
		OutFuncI:       outFunc,
		TypeTransformI: &json2struct.CppTypeTransform{},
		OutString:      []string{},
	}
	toFunc.Object([]byte("root"), data)

	h, err := util.NewFile(out.Path, out.OutDir, ".h")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer h.Close()
	header :=
		`// write by json2struct tools
#pragma once
`
	h.WriteString(header)
	h.WriteString(outFunc.Include())
	h.WriteString(fmt.Sprintf("namespace %s {\n", util.FileName(out.Path)))
	for _, item := range toStruct.OutString {
		h.WriteString(item)
	}
	for _, item := range toFunc.OutString {
		h.WriteString(item)
	}
	h.WriteString(outFunc.RootFunc())
	h.WriteString(fmt.Sprintf("\n}\n"))
}

func JsonToRapidJson(data []byte, out util.OutInfo) {
	JsonToCpp(data, out, &RapidJsonParseFuncOut{})
}
