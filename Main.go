package main

import (
	"awesomeProject/CommandLineArguments"
	"fmt"
)

func main() {
	fmt.Println("------------------------- JSON ----------------------------")
	//JSON.ParseStructJSON()
	//JSON.ParseNonStruct()
	fmt.Println("------------------------- Read File ----------------------------")
	//ReadFile.ExampleReadFile()
	//ReadFile.Read1Line()
	//ReadFile.ExampleWriteFile()
	fmt.Println("------------------------- Command Line Arguments  ----------------------------")
	//CommandLineArguments.CLAExample()
	CommandLineArguments.ParseTypeFlag()
}
