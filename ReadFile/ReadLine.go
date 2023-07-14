package ReadFile

import (
	"bufio"
	"fmt"
	"os"
)

func Read1Line() {
	fileName := "vuong1.txt"
	scanner := bufio.NewScanner(os.Stdin)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}
}
