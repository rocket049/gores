# gores

#### 项目介绍
把一个文件的内容变成一个 go 语言的变量的小工具。
用来把`gtkbuilder`、`glade`文件、模板文件、图标文件等变成 `go` 语言变量，使得编译后的程序运行时不再依赖这些外部资源。

#### 使用说明

`gores -name varName -f inputfile -pkg packageName`

```
gores -h
Usage of gores:
  -f string
    	输入文件(input file pathname)
  -name string
    	变量名称（variable name）
  -pkg string
    	package 名称(package name)，默认值：main (default "main")
```

效果： 把 'inputfile' 的内容变成变量 'varName' 的值，类型： []byte

生成： varNameRes.go

#### 安装

`go get gitee.com/rocket049/gores`
