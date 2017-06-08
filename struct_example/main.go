package main

import "fmt"

type address struct {
	address, city, state, zip string
}

func (a address) String() string {
	return a.address + "\n" + a.city + ", " + a.state + " " + a.zip + "\n"
}

type name string

func main() {
	addressMap := map[name]address{
		"Cahoon": {
			address: "4749 Raven Run",
			city:    "Broomfield",
			state:   "CO",
			zip:     "80023",
		},
		"Smith": {
			address: "4751 Raven Run",
			city:    "Broomfield",
			state:   "CO",
			zip:     "80023",
		},
	}

	fmt.Printf("%s\n%s\n", "Cahoon", addressMap["Cahoon"])
	fmt.Printf("%s\n%s\n", "Smith", addressMap["Smith"])
}
