package main

import "fmt"

type Food struct {
	ID       int
	Name     string
	Color    string
	Quantity int
}

func (f Food) String() string {
	return fmt.Sprintf(
		"Name: \t\t%q\n "+
			"Color:\t\t%q\n"+
			"Quantity:\t%v\n", f.Name, f.Color, f.Quantity)
}

var foods = []Food{
	Food{
		ID:       1,
		Name:     "banana",
		Color:    "yellow",
		Quantity: 10,
	},
	Food{
		ID:       2,
		Name:     "apple",
		Color:    "red",
		Quantity: 6,
	},
	Food{
		ID:       3,
		Name:     "onion",
		Color:    "white",
		Quantity: 8,
	},
	Food{
		ID:       4,
		Name:     "garlic",
		Color:    "purple",
		Quantity: 6,
	},
	Food{
		ID:       5,
		Name:     "carrot",
		Color:    "orange",
		Quantity: 5,
	},
	Food{
		ID:       6,
		Name:     "strawberry",
		Color:    "red",
		Quantity: 20,
	},
	Food{
		ID:       7,
		Name:     "bluberry",
		Color:    "blue",
		Quantity: 40,
	},
	Food{
		ID:       8,
		Name:     "papaya",
		Color:    "orange",
		Quantity: 3,
	},
	Food{
		ID:       9,
		Name:     "squash",
		Color:    "orange",
		Quantity: 1,
	},
	Food{
		ID:       10,
		Name:     "ginger",
		Color:    "brown",
		Quantity: 1,
	},
}
