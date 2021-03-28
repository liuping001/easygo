// Author: coolliu
// Date: 2021/3/21

package main

import (
	"fmt"
	"github.com/liuping001/easygo/draw"
	"net/http"
)

type Info struct {
	name string
	age  int
}

func main() {
	http.HandleFunc("/slice", sliceSvg)
	http.HandleFunc("/map", mapSvg)
	http.HandleFunc("/struct", structSvg)
	http.HandleFunc("/", index)

	fmt.Println("Server started at port 8888")
	http.ListenAndServe(":8888", nil)
}

func sliceSvg(w http.ResponseWriter, r *http.Request) {
	items := []string{"冰冰", "张三", "李四"}
	svg, err := draw.DrawSvg(&items)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	fmt.Fprint(w, svg)
}

func mapSvg(w http.ResponseWriter, r *http.Request) {
	item := map[string]string{}
	item["姓名"] = "冰冰"
	item["职业"] = "程序员"
	item["iphone"] = "181****9443"

	svg, err := draw.DrawSvg(&item)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	fmt.Fprint(w, svg)
}

func structSvg(w http.ResponseWriter, r *http.Request) {
	type girl struct {
		name string
		age  int
	}
	type boy struct {
		name        string
		girlFriends []girl
	}

	item := boy{
		name:        "渣男",
		girlFriends: []girl{{"琪琪", 20}, {"婷婷", 18}},
	}

	svg, err := draw.DrawSvg(&item)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	fmt.Fprint(w, svg)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<div>* <a href=\"/struct\">/struct</a></div>\n")
	fmt.Fprintf(w, "<div>* <a href=\"/map\">/map</a></div>\n")
	fmt.Fprintf(w, "<div>* <a href=\"/slice\">/slice</a></div>\n")
}
