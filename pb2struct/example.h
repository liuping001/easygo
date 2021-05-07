// write by json2struct tools
#pragma once

#include <string>
#include <vector>
#include <rapidjson/document.h>
#include <rapidjson/stringbuffer.h>
#include <rapidjson/writer.h>

namespace example {
struct Detail { 
    int64_t id = 0;
    std::string name;
    std::string age;
};
struct Data { 
    int64_t total = 0;
    std::vector<Detail> detail;
};
struct Root { 
    std::string log_id;
    bool success = false;
    Data data;
};

void ToJson(rapidjson::Writer<rapidjson::StringBuffer> &writer, const Detail &from_data_) {
    writer.StartObject();
    writer.Key("id");
    writer.Int64(from_data_.id);
    writer.Key("name");
    writer.String(from_data_.name.c_str());
    writer.Key("age");
    writer.String(from_data_.age.c_str());
    writer.EndObject();
}
void ToJson(rapidjson::Writer<rapidjson::StringBuffer> &writer, const Data &from_data_) {
    writer.StartObject();
    writer.Key("total");
    writer.Int64(from_data_.total);
    writer.Key("detail");
    writer.StartArray();
    for (const auto &item : from_data_.detail) 
    {
        ToJson(item);
    }
    writer.EndArray();
    writer.EndObject();
}
void ToJson(rapidjson::Writer<rapidjson::StringBuffer> &writer, const Root &from_data_) {
    writer.StartObject();
    writer.Key("log_id");
    writer.String(from_data_.log_id.c_str());
    writer.Key("success");
    writer.Bool(from_data_.success);
    writer.Key("data");
    ToJson(from_data_.data);
    writer.EndObject();
}
inline std::string ToJson(const Root &data) 
{
    rapidjson::StringBuffer buffer;
    rapidjson::Writer<rapidjson::StringBuffer> writer(buffer);
	ToJson(writer, data)
    return buffer.GetString();
}

}
