package main

import (
	"encoding/json"
	"fmt"
	"hellogo/example9_structwithjson/model"
	"log"
)

func main() {

	var rec model.Rectagle

	jsonBlob := []byte(`{
		"width":48,
		"length":64,
		"border":"soild"
	}`)
	err := json.Unmarshal(jsonBlob, &rec)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rec)
	fmt.Println(nil)
}
