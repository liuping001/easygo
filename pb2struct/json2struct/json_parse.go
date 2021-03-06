// Author: coolliu
// Date: 2021/4/25

package json2struct

import (
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/liuping001/easygo/pb2struct/common"
	"github.com/liuping001/easygo/pb2struct/util"
)

// value:传入数组的值，返回数组中元素的类型
func arrayItemType(value []byte) jsonparser.ValueType {
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

// 返回数组本身的类型名
func ArrayType(key []byte, value []byte, tt common.TypeTransformI) string {
	dataType := arrayItemType(value)
	if dataType == jsonparser.Object || dataType == jsonparser.Array {
		return util.CamelName(string(key))
	} else {
		return tt.StructDataType(dataType)
	}
}

// 从json解析出结构体
type JsonOutStruct struct {
	common.OutStructI
	common.TypeTransformI
	OutString []string
}

func (t *JsonOutStruct) Object(key []byte, value []byte) error {
	obj := []string{}
	obj = append(obj, t.ClassBegin(util.CamelName(string(key))))
	err := jsonparser.ObjectEach(value, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		if dataType == jsonparser.Object {
			t.Object(key, value)
			obj = append(obj, fmt.Sprintf("%s", t.Field(util.CamelName(string(key)), string(key))))
		} else if dataType == jsonparser.Array {
			err := t.Array(key, value)
			if err != nil {
				return err
			}
			obj = append(obj, fmt.Sprintf("%s", t.ArrayField(ArrayType(key, value, t), string(key))))
		} else {
			obj = append(obj, fmt.Sprintf("%s", t.Field(t.StructDataType(dataType), string(key))))
		}
		return nil
	})
	if err != nil {
		return err
	}
	obj = append(obj, t.ClassEnd())
	for _, item := range obj {
		t.OutString = append(t.OutString, item)
	}
	return nil
}

func (t *JsonOutStruct) Array(key []byte, value []byte) error {
	var err error
	count := 0
	jsonparser.ArrayEach(value, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		count++
		if count > 1 {
			return
		}
		if dataType == jsonparser.Object {
			t.Object(key, value)
		} else if dataType == jsonparser.Array {
			err = fmt.Errorf("不支持二维数组: key:%s", key)
		} else {

		}
	})
	return err
}

// 生成将结构序列化和反序列化的函数
type JsonOutParseFunc struct {
	common.OutFuncI
	common.TypeTransformI
	OutString []string
}

func (t *JsonOutParseFunc) Object(key []byte, value []byte) error {
	obj := []string{}
	obj = append(obj, t.FuncBegin(util.CamelName(string(key))))
	err := jsonparser.ObjectEach(value, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		if dataType == jsonparser.Object {
			t.Object(key, value)
			obj = append(obj, fmt.Sprintf("%s", t.FuncObjectField(util.CamelName(string(key)), string(key))))
		} else if dataType == jsonparser.Array {
			t.Array(key, value)
			itemType := arrayItemType(value)
			obj = append(obj, fmt.Sprintf("%s",
				t.FuncArrayField(t.StructDataType(itemType), string(key))))
		} else {
			obj = append(obj, fmt.Sprintf("%s", t.FuncField(t.StructDataType(dataType), string(key))))
		}
		return nil
	})
	if err != nil {
		return err
	}
	obj = append(obj, t.FuncEnd())
	for _, item := range obj {
		t.OutString = append(t.OutString, item)
	}
	return nil
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

// c++类型和json类型直接的对应关系
type CppTypeTransform struct {
}

func (ctt *CppTypeTransform) StructDataType(t interface{}) string {
	value, ok := t.(jsonparser.ValueType)
	if !ok {
		return "error_type"
	}
	if value == jsonparser.Number {
		return util.CppInt64
	} else if value == jsonparser.String {
		return util.CppString
	} else if value == jsonparser.Boolean {
		return util.CppBool
	} else if value == jsonparser.Object {
		return util.CppObject
	} else {
		return "error_type"
	}
}
