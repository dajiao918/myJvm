package main

import (
	"flag"
	"fmt"
	"os"
)

// Cmd  命令行结构体
type Cmd struct {
	helpFlag bool
	versionFlag bool
	cpOption string
	class string
	args []string
}

func cmdParse() *Cmd {
	cmd := &Cmd{}

	flag.Usage = printUsage //解析失败后的帮助信息
	flag.BoolVar(&cmd.helpFlag,"help",false,"Print help message")
	flag.BoolVar(&cmd.helpFlag,"?",false,"Print help message")
	flag.BoolVar(&cmd.versionFlag, "version",false,"Print version")
	flag.StringVar(&cmd.cpOption,"classpath","","classpath")
	flag.StringVar(&cmd.cpOption,"cp","","classpath")
	//调用此方法后会把命令行的参数-help -? -version -classpath, -cp 等后面的值解析到cmd结构体中
	flag.Parse()
	//返回解析后非键值对的参数，例如 -cp /dir  arg1 arg2->返回arg1, arg2
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}

func printUsage() {
	//os.Args[0]是程序名
	fmt.Printf("Usage: %s [-options] class [args ...]\n",os.Args[0])
}

func startJvm(cmd *Cmd) {
	fmt.Printf("classpath: %s class:%s args:%v\n",cmd.cpOption,cmd.class,cmd.args)
}