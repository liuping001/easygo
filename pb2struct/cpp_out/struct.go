// Author: coolliu
// Date: 2021/4/25

package cpp_out

import (
	"fmt"
	"github.com/liuping001/easygo/pb2struct/util"
)

// 生成c++结构体的具体方法实现
type CppStructOut struct {
}

func (c *CppStructOut) ClassBegin(Type string) string {
	return fmt.Sprintf("struct %s { \n", Type)
}
func (c *CppStructOut) ClassEnd() string {
	return fmt.Sprintf("};\n")
}

func (c *CppStructOut) ClassTab() string {
	return fmt.Sprintf("    ")
}
func (c *CppStructOut) ArrayField(Type, name string) string {
	return fmt.Sprintf("%sstd::vector<%s> %s;\n", c.ClassTab(), Type, name)
}
func (c *CppStructOut) Field(Type, name string) string {
	initValue := ""
	if Type == util.CppInt64 {
		initValue = " = 0"
	} else if Type == util.CppBool {
		initValue = " = false"
	}
	return fmt.Sprintf("%s%s %s%s;\n", c.ClassTab(), Type, name, initValue)
}
