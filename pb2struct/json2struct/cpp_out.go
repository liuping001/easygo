// Author: coolliu
// Date: 2021/4/25

package main

import (
	"fmt"
	"strings"
)

type CppStructOut struct {
}

func (c *CppStructOut) ClassBegin(name string) string {
	return fmt.Sprintf("struct %s { \n", strings.Title(name))
}
func (c *CppStructOut) ClassEnd() string {
	return fmt.Sprintf("}\n")
}

func (c *CppStructOut) ClassTab() string {
	return fmt.Sprintf("    ")
}
func (c *CppStructOut) ArrayField(Type, name string) string {
	return fmt.Sprintf("std::vector<%s> %s;\n", Type, name)
}
func (c *CppStructOut) Field(Type, name string) string {
	return fmt.Sprintf("%s %s;\n", Type, name)
}
