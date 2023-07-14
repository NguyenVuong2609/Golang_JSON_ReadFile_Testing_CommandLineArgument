package JSON

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	Name string
}

func ParseNonStruct() {
	fmt.Println("----------------------- Map ---------------------------")
	a := make(map[string]employee)
	a["1"] = employee{Name: "John"}
	j, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println(string(j))
	}

	jsonString := `{"title":"Learning JSON in Golang","author":"Lanka"}`
	var book map[string]interface{}
	err1 := json.Unmarshal([]byte(jsonString), &book)
	if err1 != nil {
		panic(err1)
	}
	fmt.Printf("%+v\n", book)
}
