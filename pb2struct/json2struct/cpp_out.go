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
	return fmt.Sprintf("%sstd::vector<%s> %s;\n", c.ClassTab(), Type, name)
}
func (c *CppStructOut) Field(Type, name string) string {
	return fmt.Sprintf("%s%s %s;\n", c.ClassTab(), Type, name)
}

type CppParseFuncOut struct {
}

func (p *CppParseFuncOut) FuncTab() string {
	return "    "
}
func (p *CppParseFuncOut) ToJsonBegin(name string) string {
	ret := `
void ToJson(rapidjson::Writer<rapidjson::StringBuffer> &writer, const %s &from_data_) {
    writer.StartObject();
`
	return fmt.Sprintf(ret, name)
}

func (p *CppParseFuncOut) ToJsonEnd() string {
	ret := `
    writer.EndObject();
}
`
	return ret
}

func (p *CppParseFuncOut) ArrayField(Type, name string, itemIsObject bool) string {
	ret := `
    writer.Key("%s");
    writer.StartArray();
    for (const auto &item : from_data_.%s) 
    {
        %s;
    }
    writer.EndArray();
`
	if itemIsObject {
		return fmt.Sprintf(ret,
			name,
			name,
			"ToJson(item)")
	} else {
		return fmt.Sprintf(ret,
			name,
			name,
			fmt.Sprintf("writer.%s(item)", Type))
	}
}

func (p *CppParseFuncOut) ObjectField(Type, name string) string {
	ret := `
    writer.Key("%s");
    ToJson(from_data_.%s);
`
	return fmt.Sprintf(ret, name,
		name)
}
func (p *CppParseFuncOut) Field(Type, name string) string {
	ret := `
    writer.Key("%s");
    writer.%s(from_data_.%s);
`
	return fmt.Sprintf(ret,
		name,
		Type, name)
}
