package main

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
	class       string
	args        []string
}

/*
cmd解析函数
*/
func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}

/*
打印提示语到控制台
*/
func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}

/*
*

	模拟开启了jvm
*/
func startJVM(cmd *Cmd) {
	fmt.Printf("classpath:%s class:%s args:%v\n",
		cmd.cpOption, cmd.class, cmd.args)
}
