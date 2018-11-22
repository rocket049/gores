# gores

#### 项目介绍
把一个文件的内容变成一个 go 语言的变量的小工具。

#### 使用说明

`gores -name varName -f inputfile -pkg packageName`

效果： 把 'inputfile' 的内容变成变量 'varName' 的值，类型： []byte

生成： varNameRes.go

#### 安装

`go get https://gitee.com/rocket049/gores`
