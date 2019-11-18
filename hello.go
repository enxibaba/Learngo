package main

import (
	"flag"
	"fmt"
	"os"
)

var name string
var cmdLine = flag.NewFlagSet("question", flag.ExitOnError)

func init() {

	//flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
	//flag.CommandLine.Usage  = func() {
	//	_, _ = fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	//	flag.PrintDefaults()
	//}

	cmdline.StringVar(&name, "name", "Elli", "The greeting object.")

	//flag.StringVar(&name, "name", "Elli", "The greeting object.")
	//or:
	//var name1 = flag.String("name", "Elli", "The greeting object.")
}

func main() {
	// 自定义命令源码文件参数说明 。
	// 自定义Usage 函数， 要在调用Parse之前
	//flag.Usage = func() {
	//	_, _ = fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	//	flag.PrintDefaults()
	//}

	//flag.Parse()
	_ = cmdLine.Parse(os.Args[1:])

	fmt.Printf("hello %s!\n", name)
}
