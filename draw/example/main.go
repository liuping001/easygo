// Author: coolliu
// Date: 2021/3/21

package main

import (
	"evoke/draw"
	"fmt"
	"net/http"
)

type Info struct {
	name string
	age int
}

func main() {
	http.HandleFunc("/slice", sliceSvg)
	http.HandleFunc("/map", mapSvg)
	http.HandleFunc("/struct", structSvg)

	fmt.Println("Server started at port 80")
	http.ListenAndServe(":80", nil)
}

func sliceSvg(w http.ResponseWriter, r *http.Request) {
	items := []string {"冰冰","张三","李四"}
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
		age int
	}
	type boy struct {
		name string
		girlFriends []girl
	}

	item := boy{
		name: "渣男",
		girlFriends: []girl{ {"琪琪", 20}, {"婷婷", 18}},
	}

	svg, err := draw.DrawSvg(&item)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	fmt.Fprint(w, svg)
}
