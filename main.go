package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	Coffee := map[string]int{
		"Latte":      55,
		"Cappuccino": 50,
		"Americano":  45,
	}
	// convert go(Map)  object to json
	bytes, _ := json.Marshal(Coffee)

	// log value
	// result: {"Americano":45,"Cappuccino":50,"Latte":55}
	fmt.Println(string(bytes))

	// log type
	// result: Type of bytes []uint8
	fmt.Printf("Type of bytes %T\n", bytes)
}
