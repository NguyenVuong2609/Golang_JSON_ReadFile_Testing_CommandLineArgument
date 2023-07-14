package JSON

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Brand string `json:"brand"`
	Name  string `json:"name"`
	Color string `json:"color"`
	Owner Person `json:"owner"`
}
type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender bool   `json:"gender"`
}

func ParseStructJSON() {
	person := Person{"Vuong", 29, true}
	car := Car{"Audi", "A8", "White", person}
	byteArray, err := json.MarshalIndent(car, "", "   ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteArray), "<---- stringify Struct to Json")

	carJson := `{"brand": "Toyota", "name": "Camry", "color": "Black"}`
	var parseCar Car
	json.Unmarshal([]byte(carJson), &parseCar)
	fmt.Printf("Brand: %s, Name: %s, Color: %s  <---- parse Json to Struct\n", parseCar.Brand, parseCar.Name, parseCar.Color)
}
