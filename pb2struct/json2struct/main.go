// Author: coolliu
// Date: 2021/4/24

package main

func main() {
	data := `
{
    "code": 0,
    "msg": "",
    "op_time": "1619242782",
    "data": {
        "base_match_m": "300",
        "oxy_match_m": "620",
        "donate_t": "30",
        "donate_m": "3000"
    },
	"listA": [
		{
        "a": "300",
        "b": ["b"],
        "c": 30
		}
	],
	"listB": ["a"]
}
`
	t := JsonOutStruct{StructOutI(&CppStructOut{})}
	t.Object([]byte("root"), []byte(data))
	{
		t := JsonOutParseFunc{FuncOutI(&CppParseFuncOut{})}
		t.Object([]byte("root"), []byte(data))
	}
}
