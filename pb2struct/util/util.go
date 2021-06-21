// Author: coolliu
// Date: 2021/5/5

package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func ReadAll(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	return ioutil.ReadAll(file)
}

func FileName(file string) string {
	base := filepath.Base(file)
	suffix := path.Ext(base)
	return strings.TrimSuffix(base, suffix)
}

func NewFile(file string, dir string, suffix string) (*os.File, error) {
	fileName := fmt.Sprintf("%s/%s%s", dir, FileName(file), suffix)
	return os.OpenFile(
		fileName,
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
}

type OutInfo struct {
	Path    string
	OutDir  string
	OutFunc int
}

var (
	CppString = "std::string"
	CppInt64  = "int64_t"
	CppBool   = "bool"
	CppObject = "object"
)
