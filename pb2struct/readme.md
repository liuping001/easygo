
# 使用
* 编译
```shell script
go build -o pb2struct main.go
```
* 将example.json转成c++代码
```shell script
./pb2struct -file=example.json -pb_type=json -to_type=rapidjson
# rapidjson 是c++代码的一种，使用的是rapidjson这个json库
```
## json2struct
根据json文本生成c++结构体。

暂不支持json的二维数组。例如：
```json
{
  "item" : [["a", "b"], ["c", "d"]]
}
```
