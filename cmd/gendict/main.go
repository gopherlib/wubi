package main

import (
	"flag"
	"os"
	"path"

	"github.com/gopherlib/wubi/dictgen"
)

var (
	// dictFile 五笔码表文件
	dictFile string
	// pkgName 目录文件所属包名
	pkgName string
	// varName 指定字典数据的变量名称，默认使用dictFile当作变量名
	varName string
	// pattern 单个五笔码在码表文件中的正则表达式
	pattern string
	// debug 开启debug日志
	debug bool
)

func init() {
	flag.StringVar(&dictFile, "dict", "", "五笔码表文件")
	flag.StringVar(&pkgName, "pkg", "dict", "生成文件所属的包名")
	flag.StringVar(&pattern, "pattern", "", "单个五笔码在码表文件中的正则表达式")
	flag.StringVar(&varName, "var", "", "指定字典数据的变量名称，默认使用dict参数当作变量名")
	flag.BoolVar(&debug, "debug", false, "开启debug日志")
	flag.Parse()
}

func main() {
	if dictFile == "" {
		panic("缺少五笔码表文件, dict 参数不能为空")
	}

	if pattern == "" {
		panic("缺少解析码表所用的正则表达式, pattern 参数不能为空")
	}

	if varName == "" {
		varName = path.Base(dictFile)
	}

	f, err := os.Open(dictFile)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = f.Close()
	}()

	gl, err := dictgen.NewLiner(pattern, debug)
	if err != nil {
		panic(err)
	}

	dict, err := gl.ParseAll(f)
	if err != nil {
		panic(err)
	}

	err = dict.Write(varName, pkgName, os.Stdout)
	if err != nil {
		panic(err)
	}
}
