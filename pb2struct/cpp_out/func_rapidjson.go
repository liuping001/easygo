// Author: coolliu
// Date: 2021/5/5

package cpp_out

import (
	"fmt"
	"github.com/liuping001/easygo/pb2struct/util"
)

type RapidJsonParseFuncOut struct {
}

func (p *RapidJsonParseFuncOut) FuncBegin(Type string) string {
	ret := `
void ToJson(rapidjson::Writer<rapidjson::StringBuffer> &writer, const %s &from_data_) {
    writer.StartObject();`
	return fmt.Sprintf(ret, Type)
}

func (p *RapidJsonParseFuncOut) FuncEnd() string {
	ret := `
    writer.EndObject();
}`
	return ret
}

func (p *RapidJsonParseFuncOut) writeFuncName(t string) string {
	if t == util.CppString {
		return "String"
	} else if t == util.CppInt64 {
		return "Int64"
	} else if t == util.CppBool {
		return "Bool"
	} else {
		return "error_field"
	}
}

func (p *RapidJsonParseFuncOut) FuncArrayField(t string, name string) string {
	ret := `
    writer.Key("%s");
    writer.StartArray();
    for (const auto &item : from_data_.%s) 
    {
        %s;
    }
    writer.EndArray();`
	if t == util.CppObject {
		return fmt.Sprintf(ret,
			name,
			name,
			"ToJson(item)")
	} else {
		suffix := ""
		if t == util.CppString {
			suffix = ".c_str()"
		}
		return fmt.Sprintf(ret,
			name,
			name,
			fmt.Sprintf("writer.%s(item%s)", p.writeFuncName(t), suffix))
	}
}

func (p *RapidJsonParseFuncOut) FuncObjectField(Type, name string) string {
	ret := `
    writer.Key("%s");
    ToJson(writer, from_data_.%s);`
	return fmt.Sprintf(ret, name,
		name)
}
func (p *RapidJsonParseFuncOut) FuncField(t, name string) string {
	suffix := ""
	if t == util.CppString {
		suffix = ".c_str()"
	}
	ret := `
    writer.Key("%s");
    writer.%s(from_data_.%s%s);`
	return fmt.Sprintf(ret, name, p.writeFuncName(t), name, suffix)
}

func (p *RapidJsonParseFuncOut) Include() string {
	return `
#include <string>
#include <vector>
#include <rapidjson/document.h>
#include <rapidjson/stringbuffer.h>
#include <rapidjson/writer.h>

`
}

func (p *RapidJsonParseFuncOut) RootFunc() string {
	return `
inline std::string ToJson(const Root &data) 
{
    rapidjson::StringBuffer buffer;
    rapidjson::Writer<rapidjson::StringBuffer> writer(buffer);
	ToJson(writer, data);
    return buffer.GetString();
}
`
}

type RapidJsonParseFuncIn struct {
}

func (p *RapidJsonParseFuncIn) FuncBegin(Type string) string {
	ret := `
void FromJson(const rapidjson::Value &doc, %s &to_data_) {
    `
	return fmt.Sprintf(ret, Type)
}

func (p *RapidJsonParseFuncIn) FuncEnd() string {
	ret := `
}`
	return ret
}

func (p *RapidJsonParseFuncIn) readFuncName(t string) string {
	if t == util.CppString {
		return "GetString"
	} else if t == util.CppInt64 {
		return "GetInt64"
	} else if t == util.CppBool {
		return "GetBool"
	} else {
		return "error_field"
	}
}

func (p *RapidJsonParseFuncIn) FuncArrayField(t string, name string) string {
	ret := `
	if (doc.HasMember("%s")) {
		auto items = doc["%s"].GetArray();
		for (auto iter = items.Begin(); iter != items.End(); iter ++)
		{
			to_data_.%s.emplace_back();
			auto &item =to_data_.%s.back();
			%s
		}
	}`
	if t == util.CppObject {
		return fmt.Sprintf(ret,
			name,
			name,
			name,
			name,
			"FromJson(*iter, item)")
	} else {
		return fmt.Sprintf(ret,
			name,
			name,
			name,
			name,
			fmt.Sprintf("item = iter->%s();", p.readFuncName(t)))
	}
}

func (p *RapidJsonParseFuncIn) FuncObjectField(Type, name string) string {
	ret := `
    if (doc.HasMember("%s")) FromJson(doc["%s"], to_data_.%s);`
	return fmt.Sprintf(ret, name, name, name)
}
func (p *RapidJsonParseFuncIn) FuncField(t, name string) string {
	ret := `
    if (doc.HasMember("%s")) to_data_.%s = doc["%s"].%s();`
	return fmt.Sprintf(ret, name, name, name, p.readFuncName(t))
}

func (p *RapidJsonParseFuncIn) Include() string {
	return `
#include <string>
#include <vector>
#include <rapidjson/document.h>
#include <rapidjson/stringbuffer.h>
#include <rapidjson/writer.h>

`
}

func (p *RapidJsonParseFuncIn) RootFunc() string {
	return `
template<class T>
T FromJson(const rapidjson::Document &doc) 
{
    T data;
	FromJson(doc, data);
    return data;
}
`
}
