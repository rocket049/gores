//gores usage: gores -name varName -f inputfile -pkg packageName
//效果： 把 'inputfile' 的内容变成变量 'varName' 的值，类型： []byte
//生成： varNameRes.go
//作者： Fu Huizhong<fuhuizn@163.com>
//Licence: GPL v2
package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

func main() {
	var name = flag.String("name", "", "变量名称（variable name）")
	var pkg = flag.String("pkg", "main", "package 名称(package name)，默认值：main")
	var fname = flag.String("f", "", "输入文件(input file pathname)")
	flag.Parse()
	tpl1 := `package %s
var %s = {{.}}
`
	if len(*name) == 0 || len(*fname) == 0 {
		panic(errors.New("name 或 f 参数未空"))
	}
	tpl := fmt.Sprintf(tpl1, *pkg, *name)
	b, err := ioutil.ReadFile(*fname)
	if err != nil {
		panic(err)
	}
	fp, err := os.Create(fmt.Sprintf("%sRes.go", *name))
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	t := template.New("")
	_, err = t.Parse(tpl)
	if err != nil {
		panic(err)
	}
	t.Execute(fp, fmt.Sprintf("%#v", b))
}
