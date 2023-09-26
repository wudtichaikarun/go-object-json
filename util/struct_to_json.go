package util

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Coffee struct {
	Menu     string
	Price    int
	Quantity int
}

type CoffeeComplex struct {
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Price       float64      `json:"price"`
	Ingredients []Ingredient `json:"ingredients"`
}

type Ingredient struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Unit     string `json:"unit"`
}

type MenuItem struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	// json:"-" = ใช้ – เพื่อทำการ ignore ฟิลด์นั้น ๆ
	// Ingredients []Ingredient `json:"-"`
	Ingredients []Ingredient `json:"ingredients"`
}

type Menu struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func ConvertStructToJsonSimple() {
	myCoffee := Coffee{"Cappuccino", 50, 2}
	// convert go(Map)  object to json
	bytes, _ := json.Marshal(myCoffee)

	// log value
	// result: {"Americano":45,"Cappuccino":50,"Latte":55}
	fmt.Println(string(bytes))

	// log type
	// result: Type of bytes []uint8
	fmt.Printf("Type of bytes %T\n", bytes)
}

func ConvertStructToJsonComplex() {
	// Define a slice of menu items
	menu := []MenuItem{
		{
			Name:        "Americano",
			Description: "A classic espresso drink",
			Price:       2.50,
			Ingredients: []Ingredient{
				{
					Name:     "Espresso",
					Quantity: 1,
					Unit:     "shot",
				},
				{
					Name:     "Water",
					Quantity: 6,
					Unit:     "oz",
				},
			},
		},
		{
			Name:        "Latte",
			Description: "Espresso with steamed milk",
			Price:       4.00,
			Ingredients: []Ingredient{
				{
					Name:     "Espresso",
					Quantity: 2,
					Unit:     "shot",
				},
				{
					Name:     "Milk",
					Quantity: 8,
					Unit:     "oz",
				},
			},
		},
	}

	// Marshal the menu items as a JSON string
	jsonString, err := json.Marshal(menu)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the JSON string to the console
	// result: [{"name":"Americano","description":"A classic espresso drink","price":2.5,"ingredients":[{"name":"Espresso","quantity":1,"unit":"shot"},{"name":"Water","quantity":6,"unit":"oz"}]},{"name":"Latte","description":"Espresso with steamed milk","price":4,"ingredients":[{"name":"Espresso","quantity":2,"unit":"shot"},{"name":"Milk","quantity":8,"unit":"oz"}]}]
	// จะออกมายาวและติดกันซึ่งดูยาก

	fmt.Println(string(jsonString))

	jsonIndentString, err := json.MarshalIndent(menu, "", " ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	/**
		[
	 {
	  "name": "Americano",
	  "description": "A classic espresso drink",
	  "price": 2.5,
	  "ingredients": [
	   {
	    "name": "Espresso",
	    "quantity": 1,
	    "unit": "shot"
	   },
	   {
	    "name": "Water",
	    "quantity": 6,
	    "unit": "oz"
	   }
	  ]
	 },
	 {
	  "name": "Latte",
	  "description": "Espresso with steamed milk",
	  "price": 4,
	  "ingredients": [
	   {
	    "name": "Espresso",
	    "quantity": 2,
	    "unit": "shot"
	   },
	   {
	    "name": "Milk",
	    "quantity": 8,
	    "unit": "oz"
	   }
	  ]
	 }
	] */
	fmt.Println(string(jsonIndentString))
}

func ConvertJsonToGoObjectSimple() {
	// Define a JSON string representing a menu item
	jsonString := `{
		"name": "Americano",
		"description": "A classic espresso drink",
		"price": 2.50
	}`

	// Unmarshal the JSON string into a Menu struct
	var menu Menu
	err := json.Unmarshal([]byte(jsonString), &menu)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the menu item to the console
	// result: Americano (A classic espresso drink): $2.50
	fmt.Printf("%s (%s): $%.2f\n", menu.Name, menu.Description, menu.Price)
}

func ConvertJsonToGoObjectComplex() {
	jsonString := `[
		{
				"name": "Americano",
				"description": "A classic espresso drink",
				"price": 2.50,
				"ingredients": [
						{
								"name": "Espresso",
								"quantity": 1,
								"unit": "shot"
						},
						{
								"name": "Water",
								"quantity": 6,
								"unit": "oz"
						}
				]
		},
		{
				"name": "Latte",
				"description": "Espresso with steamed milk",
				"price": 4.00,
				"ingredients": [
						{
								"name": "Espresso",
								"quantity": 2,
								"unit": "shot"
						},
						{
								"name": "Milk",
								"quantity": 8,
								"unit": "oz"
						}
				]
		}
	]`

	var menuItems []MenuItem
	err := json.Unmarshal([]byte(jsonString), &menuItems)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// result:
	// Americano (A classic espresso drink): $2.50
	// Latte (Espresso with steamed milk): $4.00
	for _, menuItem := range menuItems {
		fmt.Printf("%s (%s): $%.2f\n", menuItem.Name, menuItem.Description, menuItem.Price)
	}
}

func ConvertJsonFileToGoObjectComplex() {
	bytes, err := os.ReadFile("coffee.json")
	if err != nil {
		log.Fatal(err)
	}

	var coffees []CoffeeComplex
	err = json.Unmarshal(bytes, &coffees)
	if err != nil {
		log.Fatal(err)
	}

	/* result:
	Name: Americano
	Description: A classic espresso drink
	Price: 2.50
	Ingredients:
	- Espresso: 1 shot
	- Water: 6 oz

	Name: Latte
	Description: Espresso with steamed milk
	Price: 4.00
	Ingredients:
	- Espresso: 2 shot
	- Milk: 8 oz */
	for _, coffee := range coffees {
		fmt.Printf("Name: %s\n", coffee.Name)
		fmt.Printf("Description: %s\n", coffee.Description)
		fmt.Printf("Price: %.2f\n", coffee.Price)
		fmt.Println("Ingredients:")
		for _, ingredient := range coffee.Ingredients {
			fmt.Printf("- %s: %d %s\n", ingredient.Name, ingredient.Quantity, ingredient.Unit)
		}
		fmt.Println()
	}
}

func WriterGoObjectToJsonFile() {
	coffees := []CoffeeComplex{
		{
			Name:        "Americano",
			Description: "A classic espresso drink",
			Price:       2.50,
			Ingredients: []Ingredient{
				{Name: "Espresso", Quantity: 1, Unit: "shot"},
				{Name: "Water", Quantity: 6, Unit: "oz"},
			},
		},
		{
			Name:        "Latte",
			Description: "Espresso with steamed milk",
			Price:       4.00,
			Ingredients: []Ingredient{
				{Name: "Espresso", Quantity: 2, Unit: "shot"},
				{Name: "Milk", Quantity: 8, Unit: "oz"},
			},
		},
	}

	file, err := os.Create("coffee-new.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(coffees)
	if err != nil {
		log.Fatal(err)
	}
}
