package main

import (
	"fmt"
	"jvm-demo/chapter-06/classpath"
	"jvm-demo/chapter-06/rtda/heap"
	"strings"
)

func main() {
	cmd := parseCmd()

	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
		//} else if cmd.helpFlag || cmd.class == "" {
		//	printUsage()
	} else {
		//cmd.XjreOption = "/Users/hayashiarihiroshi/Library/Java/JavaVirtualMachines/corretto-1.8.0_312/Contents/Home/jre"
		cmd.class = "jvmgo.book.ch05.GaussTest"
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	classLoader := heap.NewClassLoader(cp)

	className := strings.Replace(cmd.class, ".", "/", -1)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		interpret(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}
}
