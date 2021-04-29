// Author: coolliu
// Date: 2021/4/25

package main

import (
	"fmt"
	"github.com/buger/jsonparser"
	"strings"
)

// value:传入数组的值，返回数组中元素的类型
func ArrayItemType(value []byte) jsonparser.ValueType {
	var ret jsonparser.ValueType
	retType := &ret
	count := 0
	jsonparser.ArrayEach(value, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		count++
		if count > 1 {
			return
		}
		*retType = dataType
	})
	return ret
}

type TypeTransformI interface {
	StructDataType(t jsonparser.ValueType) string
}

// 返回数组本身的类型名
func ArrayType(key []byte, value []byte, tt TypeTransformI) string {
	dataType := ArrayItemType(value)
	if dataType == jsonparser.Object || dataType == jsonparser.Array {
		return strings.Title(string(key))
	} else {
		return tt.StructDataType(dataType)
	}
}

type StructOutI interface {
	ClassBegin(name string) string
	ClassEnd() string
	ClassTab() string
	ArrayField(Type, name string) string
	Field(Type, name string) string
}

// 从json解析出结构体
type JsonOutStruct struct {
	StructOutI
	TypeTransformI
}

func (t *JsonOutStruct) Object(key []byte, value []byte) {
	obj := []string{}
	obj = append(obj, t.ClassBegin(string(key)))
	jsonparser.ObjectEach(value, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		if dataType == jsonparser.Object {
			t.Object(key, value)
			obj = append(obj, fmt.Sprintf("%s", t.Field(strings.Title(string(key)), string(key))))
		} else if dataType == jsonparser.Array {
			t.Array(key, value)
			obj = append(obj, fmt.Sprintf("%s", t.ArrayField(ArrayType(key, value, t), string(key))))
		} else {
			obj = append(obj, fmt.Sprintf("%s", t.Field(t.StructDataType(dataType), string(key))))
		}
		return nil
	})
	obj = append(obj, t.ClassEnd())
	for _, item := range obj {
		fmt.Printf(item)
	}
}

func (t *JsonOutStruct) Array(key []byte, value []byte) {
	count := 0
	jsonparser.ArrayEach(value, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		count++
		if count > 1 {
			return
		}
		if dataType == jsonparser.Object {
			t.Object(key, value)
		} else if dataType == jsonparser.Array {

		} else {

		}
	})
}

type FuncOutI interface {
	FuncTab() string
	ToJsonBegin(name string) string
	ToJsonEnd() string
	ArrayField(t jsonparser.ValueType, name string) string
	ObjectField(Type, name string) string
	Field(t jsonparser.ValueType, name string) string
}

// 从json解析出结构体
type JsonOutParseFunc struct {
	FuncOutI
	TypeTransformI
}

func (t *JsonOutParseFunc) Object(key []byte, value []byte) {
	obj := []string{}
	obj = append(obj, t.ToJsonBegin(string(key)))
	jsonparser.ObjectEach(value, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		if dataType == jsonparser.Object {
			t.Object(key, value)
			obj = append(obj, fmt.Sprintf("%s", t.ObjectField(strings.Title(string(key)), string(key))))
		} else if dataType == jsonparser.Array {
			t.Array(key, value)
			itemType := ArrayItemType(value)
			obj = append(obj, fmt.Sprintf("%s",
				t.ArrayField(itemType, string(key))))
		} else {
			obj = append(obj, fmt.Sprintf("%s", t.Field(dataType, string(key))))
		}
		return nil
	})
	obj = append(obj, t.ToJsonEnd())
	for _, item := range obj {
		fmt.Printf(item)
	}
}

func (t *JsonOutParseFunc) Array(key []byte, value []byte) {
	count := 0
	jsonparser.ArrayEach(value, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		count++
		if count > 1 {
			return
		}
		if dataType == jsonparser.Object {
			t.Object(key, value)
		} else if dataType == jsonparser.Array {

		} else {

		}
	})
}
