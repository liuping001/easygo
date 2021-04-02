// Copyright 2021
// Author: coolliu
// Date: 2021/4/2

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func ToJsonString(data interface{}) string {
	ret, _ := json.Marshal(data)
	return string(ret)
}

type UpFileRet struct {
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	FileName string `json:"file_name"`
}

func HumanSize(s int64) string {
	size := float64(s)
	if size < 1024*1024 {
		return fmt.Sprintf("%.2fk", size/1024)
	}
	return fmt.Sprintf("%.2fM", size/(1024*1024))
}

func toFile(name string, r io.Reader) error {
	file, err := os.Create(name)
	if err != nil {
		err = fmt.Errorf("os create file failed. name:%s, err:%s", name, err.Error())
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, r)
	if err != nil {
		err = fmt.Errorf("file write failed. name:%s, err:%s", name, err.Error())
		return err
	}
	return nil
}

func upFile(w http.ResponseWriter, req *http.Request) {
	ret := &UpFileRet{}
	req.ParseMultipartForm(1024 * 1024 * 100)
	mFrom := req.MultipartForm
	for k, _ := range mFrom.File {
		file, fileHeader, err := req.FormFile(k)
		if err != nil {
			ret.Msg = fmt.Sprintf("error FormFile:%s", err.Error())
			fmt.Errorf("%s\n", ret.Msg)
			fmt.Fprint(w, ToJsonString(ret))
			return
		}
		defer file.Close()
		fmt.Printf("fileName:%s, fileSize:%v\n", fileHeader.Filename, HumanSize(fileHeader.Size))

		err = toFile(fileHeader.Filename, file)
		if err != nil {
			ret.Msg = err.Error()
			fmt.Errorf("%s\n", ret.Msg)
			fmt.Fprint(w, ToJsonString(ret))
			return
		}
		ret.FileName = fileHeader.Filename
	}
	fmt.Fprint(w, ToJsonString(ret))
}

func showUpFile(w http.ResponseWriter, req *http.Request) {
	html :=
		`
<!DOCTYPE html>
<html>
<style>
#myFrom{ 
    margin:auto;
    width:70%;
    margin-top: 80px;
    margin-bottom: 50px;
}
</style>
<body>
<div id="myFrom">
	<FORM method="POST" action="/up_file" enctype="multipart/form-data" method="post">
		<INPUT type="file" name="pic"> 
		<INPUT type="submit" value="Submit">
	 </FORM>
</dev>
</body>
</html>
`
	fmt.Fprint(w, html)
}

func main() {
	http.HandleFunc("/up_file", upFile)
	http.HandleFunc("/", showUpFile)
	fmt.Println("Server started at port 9000")
	http.ListenAndServe(":9000", nil)
}
