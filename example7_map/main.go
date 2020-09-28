package main

import (
	"fmt"
)

func main() {
	n := make(map[string]string)

	n["TH"] = "THAILAND"
	n["JP"] = "JAPAN"
	n["EN"] = "ENGLISH"
	fmt.Println(n["JP"])

	y := map[string]string{
		"H":  "Hydogen",
		"Li": "Lithium",
		"B":  "Boron",
	}
	fmt.Println(y["H"])
	fmt.Println(len(y))
}
