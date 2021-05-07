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
    ToJson(from_data_.%s);`
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
	ToJson(writer, data)
    return buffer.GetString();
}
`
}
