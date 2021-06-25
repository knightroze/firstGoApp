package main

import (
	"flag"
	"fmt"
	"mypkg/link"
	"os"
	"strconv"
)

func main() {
	var inputArgv string
	flag.StringVar(&inputArgv, "input", "", "just input anything you want")

	var intInput = flag.Int("num", -1, "input a int number")
	var ynBool bool
	flag.BoolVar(&ynBool, "yn", false, "true or false")

	goCmd := flag.NewFlagSet("go", flag.ExitOnError)
	var goinput string
	goCmd.StringVar(&goinput, "input", "", "just input anything you want for go")
	goInt := goCmd.Int("int", 0, "int for go")

	goCmd2 := flag.NewFlagSet("go2", flag.ExitOnError)
	var goinput2 string
	goCmd2.StringVar(&goinput2, "input", "", "just input anything you want for go")
	goInt2 := goCmd2.Int("int", 0, "int for go")

	flag.Parse()
	fmt.Println(flag.NArg())
	fmt.Println(flag.NFlag())

	switch {
	case inputArgv != "":
		{
			fmt.Println("inputArgv is ", inputArgv)

		}
		fallthrough //不能包含在case里面
	case *intInput != -1:
		{
			fmt.Println("intInput is ", strconv.Itoa(*intInput))
		}

	}
	if flag.Lookup("input1") == nil {
		fmt.Println("input has not been input ")
	}

	fmt.Println(flag.Args())
	goCmd.Parse(flag.Args()[2:])
	goCmd2.Parse(flag.Args()[4:])
	fmt.Println(*goInt)
	fmt.Println(*goInt2)
	flag.PrintDefaults()
	fmt.Println(os.Args)
	link.PrintTest("hello world for link")
	link.PrintTest("sd")
}
