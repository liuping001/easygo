// Author: coolliu
// Date: 2021/5/5

package common

// 输出结构体
type OutStructI interface {
	ClassBegin(Type string) string
	ClassEnd() string
	ClassTab() string
	ArrayField(Type, name string) string
	Field(Type, name string) string
}

// 输出序列化，反序列化函数
type OutFuncI interface {
	FuncBegin(Type string) string
	FuncEnd() string
	FuncArrayField(Type, name string) string
	FuncObjectField(Type, name string) string
	FuncField(Type, name string) string
}

// pb -> 输出语言 的type转化
// 如 json的number -> cpp的int64_t
type TypeTransformI interface {
	StructDataType(t interface{}) string
}
