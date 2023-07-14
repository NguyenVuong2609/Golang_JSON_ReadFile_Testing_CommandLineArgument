package CommandLineArguments

import (
	"fmt"
	"os"
)

func CLAExample() {
	// The first argument
	// is always program name
	myProgramName := os.Args[0]
	cmdArgs := os.Args[4]
	gettingArgs := os.Args[2]
	toGetAllArgs := os.Args[1:]

	// it will display
	// the program name
	fmt.Println(myProgramName)
	fmt.Println(cmdArgs)
	fmt.Println(gettingArgs)
	fmt.Println(toGetAllArgs)
	for n, args := range os.Args {
		fmt.Println("Arg", n, "->", args)
	}
}
