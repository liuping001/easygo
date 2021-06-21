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

func JsonToCpp(data []byte, out util.OutInfo, toJson, fromJson cppOutFuncI) error {
	toStruct := json2struct.JsonOutStruct{
		OutStructI:     &CppStructOut{},
		TypeTransformI: &json2struct.CppTypeTransform{},
		OutString:      []string{},
	}
	err := toStruct.Object([]byte("root"), data)
	if err != nil {
		return err
	}

	ToJsonSting := json2struct.JsonOutParseFunc{
		OutFuncI:       toJson,
		TypeTransformI: &json2struct.CppTypeTransform{},
		OutString:      []string{},
	}
	if toJson != nil {
		err = ToJsonSting.Object([]byte("root"), data)
		if err != nil {
			return err
		}
	}

	FromJsonString := json2struct.JsonOutParseFunc{
		OutFuncI:       fromJson,
		TypeTransformI: &json2struct.CppTypeTransform{},
		OutString:      []string{},
	}
	if fromJson != nil {
		err = FromJsonString.Object([]byte("root"), data)
		if err != nil {
			return err
		}
	}

	h, err := util.NewFile(out.Path, out.OutDir, ".h")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer h.Close()
	header :=
		`// write by json2struct tools
#pragma once
`
	h.WriteString(header)

	if toJson != nil {
		h.WriteString(toJson.Include())
	} else {
		h.WriteString(fromJson.Include())
	}

	h.WriteString(fmt.Sprintf("namespace %s {\n", util.FileName(out.Path)))
	for _, item := range toStruct.OutString {
		h.WriteString(item)
	}
	if toJson != nil {
		for _, item := range ToJsonSting.OutString {
			h.WriteString(item)
		}
		h.WriteString(toJson.RootFunc())
	}

	if fromJson != nil {
		for _, item := range FromJsonString.OutString {
			h.WriteString(item)
		}
		h.WriteString(fromJson.RootFunc())
	}

	h.WriteString(fmt.Sprintf("\n}\n"))
	return nil
}

func JsonToRapidJson(data []byte, out util.OutInfo) error {
	if out.OutFunc == 1 {
		return JsonToCpp(data, out, nil, &RapidJsonParseFuncIn{})
	} else if out.OutFunc == 2 {
		return JsonToCpp(data, out, &RapidJsonParseFuncOut{}, nil)
	} else {
		return JsonToCpp(data, out, &RapidJsonParseFuncOut{}, &RapidJsonParseFuncIn{})
	}
}
