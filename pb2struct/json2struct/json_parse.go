// Author: coolliu
// Date: 2021/4/25

package main

import (
	"fmt"
	"github.com/buger/jsonparser"
	"strings"
)

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
			obj = append(obj, fmt.Sprintf("%s", t.ArrayField(t.ArrayType(key, value), string(key))))
		} else {
			obj = append(obj, fmt.Sprintf("%s", t.Field(dataType.String(), string(key))))
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

func (t *JsonOutStruct) ArrayType(key []byte, value []byte) string {
	var ret string
	retType := &ret
	count := 0
	jsonparser.ArrayEach(value, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		count++
		if count > 1 {
			return
		}
		if dataType == jsonparser.Object || dataType == jsonparser.Array {
			*retType = strings.Title(string(key))
		} else {
			*retType = dataType.String()
		}
	})
	return ret
}

type FuncOutI interface {
	FuncTab() string
	ToJsonBegin(name string) string
	ToJsonEnd() string
	ArrayField(Type, name string, itemIsObject bool) string
	ObjectField(Type, name string) string
	Field(Type, name string) string
}

// 从json解析出结构体
type JsonOutParseFunc struct {
	FuncOutI
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
			itemType := t.ArrayItemType(value)
			obj = append(obj, fmt.Sprintf("%s",
				t.ArrayField(t.ArrayType(key, value), string(key), itemType == jsonparser.Object)))
		} else {
			obj = append(obj, fmt.Sprintf("%s", t.Field(dataType.String(), string(key))))
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

func (t *JsonOutParseFunc) ArrayItemType(value []byte) jsonparser.ValueType {
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

func (t *JsonOutParseFunc) ArrayType(key []byte, value []byte) string {
	var ret string
	retType := &ret
	count := 0
	jsonparser.ArrayEach(value, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		count++
		if count > 1 {
			return
		}
		if dataType == jsonparser.Object || dataType == jsonparser.Array {
			*retType = strings.Title(string(key))
		} else {
			*retType = dataType.String()
		}
	})
	return ret
}
