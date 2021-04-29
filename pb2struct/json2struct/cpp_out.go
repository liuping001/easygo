// Author: coolliu
// Date: 2021/4/25

package main

import (
	"fmt"
	"github.com/buger/jsonparser"
	"strings"
)

type CppTypeTransform struct {
}

func (ctt *CppTypeTransform) StructDataType(t jsonparser.ValueType) string {
	if t == jsonparser.Number {
		return "int64_t"
	} else if t == jsonparser.String {
		return "std::string"
	} else if t == jsonparser.Boolean {
		return "bool"
	} else {
		return "error_type"
	}
}

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
    writer.StartObject();`
	return fmt.Sprintf(ret, name)
}

func (p *CppParseFuncOut) ToJsonEnd() string {
	ret := `
    writer.EndObject();
}`
	return ret
}

func (p *CppParseFuncOut) WriteFuncName(t jsonparser.ValueType) string {
	if t == jsonparser.String {
		return "String"
	} else if t == jsonparser.Number {
		return "Int64"
	} else if t == jsonparser.Boolean {
		return "Bool"
	} else {
		return "error_field"
	}
}

func (p *CppParseFuncOut) ArrayField(t jsonparser.ValueType, name string) string {
	ret := `
    writer.Key("%s");
    writer.StartArray();
    for (const auto &item : from_data_.%s) 
    {
        %s;
    }
    writer.EndArray();`
	if t == jsonparser.Object {
		return fmt.Sprintf(ret,
			name,
			name,
			"ToJson(item)")
	} else {
		return fmt.Sprintf(ret,
			name,
			name,
			fmt.Sprintf("writer.%s(item)", p.WriteFuncName(t)))
	}
}

func (p *CppParseFuncOut) ObjectField(Type, name string) string {
	ret := `
    writer.Key("%s");
    ToJson(from_data_.%s);`
	return fmt.Sprintf(ret, name,
		name)
}
func (p *CppParseFuncOut) Field(t jsonparser.ValueType, name string) string {
	ret := `
    writer.Key("%s");
    writer.%s(from_data_.%s);`
	return fmt.Sprintf(ret, name, p.WriteFuncName(t), name)
}
