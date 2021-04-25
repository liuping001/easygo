// Author: coolliu
// Date: 2021/4/24

package main

import (
	"fmt"
	"github.com/buger/jsonparser"
	"strings"
)

type JsonTree struct {
}

func (t *JsonTree) Object(key []byte, value []byte) {
	obj := []string{}
	obj = append(obj, fmt.Sprintf("struct %s {\n", strings.Title(string(key))))
	jsonparser.ObjectEach(value, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		if dataType == jsonparser.Object {
			t.Object(key, value)
			obj = append(obj, fmt.Sprintf("\t%s %s;\n", strings.Title(string(key)), string(key)))
		} else if dataType == jsonparser.Array {
			t.Array(key, value)
			obj = append(obj, fmt.Sprintf("\tstd::vector<%s> %s;\n", t.ArrayType(key, value), string(key)))
		} else {
			obj = append(obj, fmt.Sprintf("\t%s %s;\n", dataType.String(), string(key)))
		}
		return nil
	})
	obj = append(obj, fmt.Sprintf("}\n"))
	for _, item := range obj {
		fmt.Printf(item)
	}
}

func (t *JsonTree) Array(key []byte, value []byte) {
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

func (t *JsonTree) ArrayType(key []byte, value []byte) string {
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

func main() {
	data := `
{
    "code": 0,
    "msg": "",
    "op_time": "1619242782",
    "data": {
        "base_match_m": "300",
        "oxy_match_m": "620",
        "donate_t": "30",
        "donate_m": "3000"
    },
	"listA": [
		{
        "a": "300",
        "b": ["b"],
        "c": 30
		}
	],
	"listB": ["a"]
}
`
	var t JsonTree
	t.Object([]byte("root"), []byte(data))
}
