
# pb2struct
将json、yml、toml等文件，转成c++、go语言的代码
* 编译
```shell script
go build -o pb2struct main.go
```



## json2struct
根据json文本生成指定语言的代码。

暂不支持json的二维数组。例如：
```json
{
  "item" : [["a", "b"], ["c", "d"]]
}
```
#### 测试
* 例：test.json转成c++代码
```shell script
./pb2struct -file=test.json -pb_type=json -to_type=rapidjson
# rapidjson 是c++代码的一种，使用的是rapidjson这个json库
```

```shell script
g++ -std=c++11 -o test test.cpp
```