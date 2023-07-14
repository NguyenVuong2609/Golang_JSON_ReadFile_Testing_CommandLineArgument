package ReadFile

import (
	"fmt"
	"log"
	"os"
)

func ExampleReadFile() {
	body, err := os.ReadFile("vuong1.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	fmt.Println(string(body))

}
