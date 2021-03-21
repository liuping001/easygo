// Author: coolliu
// Date: 2021/3/21

package draw

import (
	"bytes"
	"github.com/bradleyjkemp/memviz"
	"github.com/goccy/go-graphviz"
)

// 将go的对象转换成dot
func Draw(data interface{}) string {
	buf := &bytes.Buffer{}
	memviz.Map(buf, data)
	return buf.String()
}

// 将go对象直接转svg
func DrawSvg(data interface{}) (string, error) {
	dot := Draw(data)
	return Dot2Svg(dot)
}

// dot2svg
func Dot2Svg(dot string) (string, error) {
	graph, err:=graphviz.ParseBytes([]byte(dot))
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	g := graphviz.New()
	err = g.Render(graph, graphviz.SVG, &buf)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
