package ReadFile

import (
	"fmt"
	"os"
)

type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender bool   `json:"gender"`
}

func ExampleWriteFile() {
	//person := Person{"Thu", 25, false}
	filename := "vuong.txt"
	//if _, err := os.Stat(filename); os.IsNotExist(err) {
	//	file, err := os.Create(filename)
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//	fmt.Println("File created successfully.")
	//	defer file.Close()
	//} else {
	//	fmt.Println("File exists")
	//}
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	//byteArray, err := json.Marshal(person)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//content := string(byteArray)
	//if _, err := fmt.Fprintln(file, content); err != nil {
	//	fmt.Println(err)
	//	return
	//}
}
