// write by json2struct tools
#pragma once

#include <string>
#include <vector>
#include <rapidjson/document.h>
#include <rapidjson/stringbuffer.h>
#include <rapidjson/writer.h>

namespace test {
struct MsgRedis { 
    int64_t qnum = 0;
    int64_t mod = 0;
    int64_t cmd = 0;
    std::string node;
    std::string passwd;
    int64_t timeout = 0;
};
struct DataRedis { 
    int64_t qnum = 0;
    int64_t mod = 0;
    int64_t cmd = 0;
    std::string node;
    std::string passwd;
    int64_t timeout = 0;
};
struct MatchReqRedis { 
    int64_t qnum = 0;
    int64_t mod = 0;
    int64_t cmd = 0;
    std::string node;
    std::string passwd;
    int64_t timeout = 0;
};
struct MatchResultRedis { 
    int64_t qnum = 0;
    int64_t mod = 0;
    int64_t cmd = 0;
    std::string node;
    std::string passwd;
    int64_t timeout = 0;
};
struct QualifyService { 
    int64_t thread_num = 0;
};
struct CalcService { 
    int64_t as_thread = 0;
    int64_t dm_thread = 0;
};
struct PayService { 
    int64_t thread_num = 0;
};
struct ResultService { 
    int64_t thread_num = 0;
};
struct AssistService { 
    int64_t port = 0;
    int64_t work_num = 0;
    int64_t io_num = 0;
};
struct TxBaseMatch { 
    bool use_plat = false;
    int64_t plat_mod = 0;
    int64_t plat_cmd = 0;
    std::string plat_node;
    int64_t plat_timeout = 0;
    int64_t evid = 0;
    int64_t cache_match_end_cycle = 0;
};
struct Mysql { 
    std::string ip;
    int64_t port = 0;
    std::string user;
    std::string password;
    std::string database;
    int64_t timeout = 0;
    std::string charset;
    std::string tablename;
};
struct Root { 
    std::string loglevel;
    int64_t logfile_size = 0;
    int64_t logfile_backup_num = 0;
    MsgRedis msg_redis;
    DataRedis data_redis;
    MatchReqRedis match_req_redis;
    MatchResultRedis match_result_redis;
    QualifyService qualify_service;
    CalcService calc_service;
    PayService pay_service;
    ResultService result_service;
    AssistService assist_service;
    TxBaseMatch tx_base_match;
    std::vector<std::string> qualify_flow;
    std::vector<std::string> moni_qualify_flow;
    std::vector<std::string> zl_host_flow;
    std::vector<std::string> zl_guest_flow;
    Mysql mysql;
};

void ToJson(rapidjson::Writer<rapidjson::StringBuffer> &writer, const MsgRedis &from_data_) {
    writer.StartObject();
    writer.Key("qnum");
    writer.Int64(from_data_.qnum);
    writer.Key("mod");
    writer.Int64(from_data_.mod);
    writer.Key("cmd");
    writer.Int64(from_data_.cmd);
    writer.Key("node");
    writer.String(from_data_.node.c_str());
    writer.Key("passwd");
    writer.String(from_data_.passwd.c_str());
    writer.Key("timeout");
    writer.Int64(from_data_.timeout);
    writer.EndObject();
}
void ToJson(rapidjson::Writer<rapidjson::StringBuffer> &writer, const DataRedis &from_data_) {
    writer.StartObject();
    writer.Key("qnum");
    writer.Int64(from_data_.qnum);
    writer.Key("mod");
    writer.Int64(from_data_.mod);
    writer.Key("cmd");
    writer.Int64(from_data_.cmd);
    writer.Key("node");
    writer.String(from_data_.node.c_str());
    writer.Key("passwd");
    writer.String(from_data_.passwd.c_str());
    writer.Key("timeout");
    writer.Int64(from_data_.timeout);
    writer.EndObject();
}
void ToJson(rapidjson::Writer<rapidjson::StringBuffer> &writer, const MatchReqRedis &from_data_) {
    writer.StartObject();
    writer.Key("qnum");
    writer.Int64(from_data_.qnum);
    writer.Key("mod");
    writer.Int64(from_data_.mod);
    writer.Key("cmd");
    writer.Int64(from_data_.cmd);
    writer.Key("node");
    writer.String(from_data_.node.c_str());
    writer.Key("passwd");
    writer.String(from_data_.passwd.c_str());
    writer.Key("timeout");
    writer.Int64(from_data_.timeout);
    writer.EndObject();
}
void ToJson(rapidjson::Writer<rapidjson::StringBuffer> &writer, const MatchResultRedis &from_data_) {
    writer.StartObject();
    writer.Key("qnum");
    writer.Int64(from_data_.qnum);
    writer.Key("mod");
    writer.Int64(from_data_.mod);
    writer.Key("cmd");
    writer.Int64(from_data_.cmd);
    writer.Key("node");
    writer.String(from_data_.node.c_str());
    writer.Key("passwd");
    writer.String(from_data_.passwd.c_str());
    writer.Key("timeout");
    writer.Int64(from_data_.timeout);
    writer.EndObject();
}
void ToJson(rapidjson::Writer<rapidjson::StringBuffer> &writer, const QualifyService &from_data_) {
    writer.StartObject();
    writer.Key("thread_num");
    writer.Int64(from_data_.thread_num);
    writer.EndObject();
}
void ToJson(rapidjson::Writer<rapidjson::StringBuffer> &writer, const CalcService &from_data_) {
    writer.StartObject();
    writer.Key("as_thread");
    writer.Int64(from_data_.as_thread);
    writer.Key("dm_thread");
    writer.Int64(from_data_.dm_thread);
    writer.EndObject();
}
void ToJson(rapidjson::Writer<rapidjson::StringBuffer> &writer, const PayService &from_data_) {
    writer.StartObject();
    writer.Key("thread_num");
    writer.Int64(from_data_.thread_num);
    writer.EndObject();
}
void ToJson(rapidjson::Writer<rapidjson::StringBuffer> &writer, const ResultService &from_data_) {
    writer.StartObject();
    writer.Key("thread_num");
    writer.Int64(from_data_.thread_num);
    writer.EndObject();
}
void ToJson(rapidjson::Writer<rapidjson::StringBuffer> &writer, const AssistService &from_data_) {
    writer.StartObject();
    writer.Key("port");
    writer.Int64(from_data_.port);
    writer.Key("work_num");
    writer.Int64(from_data_.work_num);
    writer.Key("io_num");
    writer.Int64(from_data_.io_num);
    writer.EndObject();
}
void ToJson(rapidjson::Writer<rapidjson::StringBuffer> &writer, const TxBaseMatch &from_data_) {
    writer.StartObject();
    writer.Key("use_plat");
    writer.Bool(from_data_.use_plat);
    writer.Key("plat_mod");
    writer.Int64(from_data_.plat_mod);
    writer.Key("plat_cmd");
    writer.Int64(from_data_.plat_cmd);
    writer.Key("plat_node");
    writer.String(from_data_.plat_node.c_str());
    writer.Key("plat_timeout");
    writer.Int64(from_data_.plat_timeout);
    writer.Key("evid");
    writer.Int64(from_data_.evid);
    writer.Key("cache_match_end_cycle");
    writer.Int64(from_data_.cache_match_end_cycle);
    writer.EndObject();
}
void ToJson(rapidjson::Writer<rapidjson::StringBuffer> &writer, const Mysql &from_data_) {
    writer.StartObject();
    writer.Key("ip");
    writer.String(from_data_.ip.c_str());
    writer.Key("port");
    writer.Int64(from_data_.port);
    writer.Key("user");
    writer.String(from_data_.user.c_str());
    writer.Key("password");
    writer.String(from_data_.password.c_str());
    writer.Key("database");
    writer.String(from_data_.database.c_str());
    writer.Key("timeout");
    writer.Int64(from_data_.timeout);
    writer.Key("charset");
    writer.String(from_data_.charset.c_str());
    writer.Key("tablename");
    writer.String(from_data_.tablename.c_str());
    writer.EndObject();
}
void ToJson(rapidjson::Writer<rapidjson::StringBuffer> &writer, const Root &from_data_) {
    writer.StartObject();
    writer.Key("loglevel");
    writer.String(from_data_.loglevel.c_str());
    writer.Key("logfile_size");
    writer.Int64(from_data_.logfile_size);
    writer.Key("logfile_backup_num");
    writer.Int64(from_data_.logfile_backup_num);
    writer.Key("msg_redis");
    ToJson(writer, from_data_.msg_redis);
    writer.Key("data_redis");
    ToJson(writer, from_data_.data_redis);
    writer.Key("match_req_redis");
    ToJson(writer, from_data_.match_req_redis);
    writer.Key("match_result_redis");
    ToJson(writer, from_data_.match_result_redis);
    writer.Key("qualify_service");
    ToJson(writer, from_data_.qualify_service);
    writer.Key("calc_service");
    ToJson(writer, from_data_.calc_service);
    writer.Key("pay_service");
    ToJson(writer, from_data_.pay_service);
    writer.Key("result_service");
    ToJson(writer, from_data_.result_service);
    writer.Key("assist_service");
    ToJson(writer, from_data_.assist_service);
    writer.Key("tx_base_match");
    ToJson(writer, from_data_.tx_base_match);
    writer.Key("qualify_flow");
    writer.StartArray();
    for (const auto &item : from_data_.qualify_flow) 
    {
        writer.String(item.c_str());
    }
    writer.EndArray();
    writer.Key("moni_qualify_flow");
    writer.StartArray();
    for (const auto &item : from_data_.moni_qualify_flow) 
    {
        writer.String(item.c_str());
    }
    writer.EndArray();
    writer.Key("zl_host_flow");
    writer.StartArray();
    for (const auto &item : from_data_.zl_host_flow) 
    {
        writer.String(item.c_str());
    }
    writer.EndArray();
    writer.Key("zl_guest_flow");
    writer.StartArray();
    for (const auto &item : from_data_.zl_guest_flow) 
    {
        writer.String(item.c_str());
    }
    writer.EndArray();
    writer.Key("mysql");
    ToJson(writer, from_data_.mysql);
    writer.EndObject();
}
inline std::string ToJson(const Root &data) 
{
    rapidjson::StringBuffer buffer;
    rapidjson::Writer<rapidjson::StringBuffer> writer(buffer);
    ToJson(writer, data);
    return buffer.GetString();
}

void FromJson(const rapidjson::Value &doc, MsgRedis &to_data_) {
    
    if (doc.HasMember("qnum")) to_data_.qnum = doc["qnum"].GetInt64();
    if (doc.HasMember("mod")) to_data_.mod = doc["mod"].GetInt64();
    if (doc.HasMember("cmd")) to_data_.cmd = doc["cmd"].GetInt64();
    if (doc.HasMember("node")) to_data_.node = doc["node"].GetString();
    if (doc.HasMember("passwd")) to_data_.passwd = doc["passwd"].GetString();
    if (doc.HasMember("timeout")) to_data_.timeout = doc["timeout"].GetInt64();
}
void FromJson(const rapidjson::Value &doc, DataRedis &to_data_) {
    
    if (doc.HasMember("qnum")) to_data_.qnum = doc["qnum"].GetInt64();
    if (doc.HasMember("mod")) to_data_.mod = doc["mod"].GetInt64();
    if (doc.HasMember("cmd")) to_data_.cmd = doc["cmd"].GetInt64();
    if (doc.HasMember("node")) to_data_.node = doc["node"].GetString();
    if (doc.HasMember("passwd")) to_data_.passwd = doc["passwd"].GetString();
    if (doc.HasMember("timeout")) to_data_.timeout = doc["timeout"].GetInt64();
}
void FromJson(const rapidjson::Value &doc, MatchReqRedis &to_data_) {
    
    if (doc.HasMember("qnum")) to_data_.qnum = doc["qnum"].GetInt64();
    if (doc.HasMember("mod")) to_data_.mod = doc["mod"].GetInt64();
    if (doc.HasMember("cmd")) to_data_.cmd = doc["cmd"].GetInt64();
    if (doc.HasMember("node")) to_data_.node = doc["node"].GetString();
    if (doc.HasMember("passwd")) to_data_.passwd = doc["passwd"].GetString();
    if (doc.HasMember("timeout")) to_data_.timeout = doc["timeout"].GetInt64();
}
void FromJson(const rapidjson::Value &doc, MatchResultRedis &to_data_) {
    
    if (doc.HasMember("qnum")) to_data_.qnum = doc["qnum"].GetInt64();
    if (doc.HasMember("mod")) to_data_.mod = doc["mod"].GetInt64();
    if (doc.HasMember("cmd")) to_data_.cmd = doc["cmd"].GetInt64();
    if (doc.HasMember("node")) to_data_.node = doc["node"].GetString();
    if (doc.HasMember("passwd")) to_data_.passwd = doc["passwd"].GetString();
    if (doc.HasMember("timeout")) to_data_.timeout = doc["timeout"].GetInt64();
}
void FromJson(const rapidjson::Value &doc, QualifyService &to_data_) {
    
    if (doc.HasMember("thread_num")) to_data_.thread_num = doc["thread_num"].GetInt64();
}
void FromJson(const rapidjson::Value &doc, CalcService &to_data_) {
    
    if (doc.HasMember("as_thread")) to_data_.as_thread = doc["as_thread"].GetInt64();
    if (doc.HasMember("dm_thread")) to_data_.dm_thread = doc["dm_thread"].GetInt64();
}
void FromJson(const rapidjson::Value &doc, PayService &to_data_) {
    
    if (doc.HasMember("thread_num")) to_data_.thread_num = doc["thread_num"].GetInt64();
}
void FromJson(const rapidjson::Value &doc, ResultService &to_data_) {
    
    if (doc.HasMember("thread_num")) to_data_.thread_num = doc["thread_num"].GetInt64();
}
void FromJson(const rapidjson::Value &doc, AssistService &to_data_) {
    
    if (doc.HasMember("port")) to_data_.port = doc["port"].GetInt64();
    if (doc.HasMember("work_num")) to_data_.work_num = doc["work_num"].GetInt64();
    if (doc.HasMember("io_num")) to_data_.io_num = doc["io_num"].GetInt64();
}
void FromJson(const rapidjson::Value &doc, TxBaseMatch &to_data_) {
    
    if (doc.HasMember("use_plat")) to_data_.use_plat = doc["use_plat"].GetBool();
    if (doc.HasMember("plat_mod")) to_data_.plat_mod = doc["plat_mod"].GetInt64();
    if (doc.HasMember("plat_cmd")) to_data_.plat_cmd = doc["plat_cmd"].GetInt64();
    if (doc.HasMember("plat_node")) to_data_.plat_node = doc["plat_node"].GetString();
    if (doc.HasMember("plat_timeout")) to_data_.plat_timeout = doc["plat_timeout"].GetInt64();
    if (doc.HasMember("evid")) to_data_.evid = doc["evid"].GetInt64();
    if (doc.HasMember("cache_match_end_cycle")) to_data_.cache_match_end_cycle = doc["cache_match_end_cycle"].GetInt64();
}
void FromJson(const rapidjson::Value &doc, Mysql &to_data_) {
    
    if (doc.HasMember("ip")) to_data_.ip = doc["ip"].GetString();
    if (doc.HasMember("port")) to_data_.port = doc["port"].GetInt64();
    if (doc.HasMember("user")) to_data_.user = doc["user"].GetString();
    if (doc.HasMember("password")) to_data_.password = doc["password"].GetString();
    if (doc.HasMember("database")) to_data_.database = doc["database"].GetString();
    if (doc.HasMember("timeout")) to_data_.timeout = doc["timeout"].GetInt64();
    if (doc.HasMember("charset")) to_data_.charset = doc["charset"].GetString();
    if (doc.HasMember("tablename")) to_data_.tablename = doc["tablename"].GetString();
}
void FromJson(const rapidjson::Value &doc, Root &to_data_) {
    
    if (doc.HasMember("loglevel")) to_data_.loglevel = doc["loglevel"].GetString();
    if (doc.HasMember("logfile_size")) to_data_.logfile_size = doc["logfile_size"].GetInt64();
    if (doc.HasMember("logfile_backup_num")) to_data_.logfile_backup_num = doc["logfile_backup_num"].GetInt64();
    if (doc.HasMember("msg_redis")) FromJson(doc["msg_redis"], to_data_.msg_redis);
    if (doc.HasMember("data_redis")) FromJson(doc["data_redis"], to_data_.data_redis);
    if (doc.HasMember("match_req_redis")) FromJson(doc["match_req_redis"], to_data_.match_req_redis);
    if (doc.HasMember("match_result_redis")) FromJson(doc["match_result_redis"], to_data_.match_result_redis);
    if (doc.HasMember("qualify_service")) FromJson(doc["qualify_service"], to_data_.qualify_service);
    if (doc.HasMember("calc_service")) FromJson(doc["calc_service"], to_data_.calc_service);
    if (doc.HasMember("pay_service")) FromJson(doc["pay_service"], to_data_.pay_service);
    if (doc.HasMember("result_service")) FromJson(doc["result_service"], to_data_.result_service);
    if (doc.HasMember("assist_service")) FromJson(doc["assist_service"], to_data_.assist_service);
    if (doc.HasMember("tx_base_match")) FromJson(doc["tx_base_match"], to_data_.tx_base_match);
    if (doc.HasMember("qualify_flow")) {
        auto items = doc["qualify_flow"].GetArray();
        for (auto iter = items.Begin(); iter != items.End(); iter ++)
        {
            to_data_.qualify_flow.emplace_back();
            auto &item =to_data_.qualify_flow.back();
            item = iter->GetString();
        }
    }
    if (doc.HasMember("moni_qualify_flow")) {
        auto items = doc["moni_qualify_flow"].GetArray();
        for (auto iter = items.Begin(); iter != items.End(); iter ++)
        {
            to_data_.moni_qualify_flow.emplace_back();
            auto &item =to_data_.moni_qualify_flow.back();
            item = iter->GetString();
        }
    }
    if (doc.HasMember("zl_host_flow")) {
        auto items = doc["zl_host_flow"].GetArray();
        for (auto iter = items.Begin(); iter != items.End(); iter ++)
        {
            to_data_.zl_host_flow.emplace_back();
            auto &item =to_data_.zl_host_flow.back();
            item = iter->GetString();
        }
    }
    if (doc.HasMember("zl_guest_flow")) {
        auto items = doc["zl_guest_flow"].GetArray();
        for (auto iter = items.Begin(); iter != items.End(); iter ++)
        {
            to_data_.zl_guest_flow.emplace_back();
            auto &item =to_data_.zl_guest_flow.back();
            item = iter->GetString();
        }
    }
    if (doc.HasMember("mysql")) FromJson(doc["mysql"], to_data_.mysql);
}
Root FromJson(const rapidjson::Document &doc) 
{
    Root data;
    FromJson(doc, data);
    return data;
}

}
