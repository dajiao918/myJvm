package main

import (
	"fmt"
	"jvmgo/ch02/classpath"
	"jvmgo/ch03/classFile"
	"strings"
)

func main() {
	cmd := cmdParse()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJvm(cmd)
	}
}

func startJvm(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath: %v class:%v args:%v\n", cp, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)
	cf := loadClass(className, cp)
	printClassInfo(cf)
	mainMethod := getMainMethod(cf)
	if mainMethod != nil {
		interpret(mainMethod)
	} else {
		fmt.Printf("Main method is nil, occur error!!!")
	}
}

func getMainMethod(cf *classFile.ClassFile) *classFile.MemberInfo {
	for _, mem := range cf.Methods() {
		fmt.Printf("name=%v, descriptor=%v\n", mem.Name(), mem.Descriptor())
		if mem.Name() == "main" && mem.Descriptor() == "([Ljava/lang/String;)V" {
			return mem
		}
	}
	return nil
}

func loadClass(name string, cp *classpath.Classpath) *classFile.ClassFile {
	classData, _, err := cp.ReadClass(name)
	if err != nil {
		panic(err)
	}
	classFile, err := classFile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return classFile
}

func printClassInfo(cf *classFile.ClassFile) {
	fmt.Printf("Version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constants count: %v\n", len(cf.ConstantPool()))
	fmt.Printf("access flags: 0x%x\n", cf.AccessFlags())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("supre class: %v\n", cf.SuperClassName())
	fmt.Printf("intefaces Name: %v\n", cf.InterfacesName())
	fmt.Printf("fields count: %v\n", len(cf.Fields()))
	for i, v := range cf.Fields() {
		fmt.Printf("field%d name is %v\n", i, v.Name())
	}
	fmt.Printf("methods count: %v\n", len(cf.Methods()))
	for i, v := range cf.Methods() {
		fmt.Printf("method%d name is %v\n", i, v.Name())
	}
}
