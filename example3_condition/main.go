package main

import (
	"fmt"
)

func main() {
	/*conditions := true

	if conditions {
		fmt.Println("A")
	}

	if conditions == true {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	if conditions {
		if conditions != true {
			fmt.Println("D")
		} else {
			fmt.Println("E")
			if v := false; v == conditions {
				fmt.Println("F")
			}
		}
	}**/

	var score int
	fmt.Print("Enter Score:")
	fmt.Scanf("%d", &score)

	if score > 80 {
		fmt.Println("A")
	} else if score <= 79 && score >= 60 {
		fmt.Println("B")
	} else if score <= 69 && score >= 50 {
		fmt.Println("C")
	} else {
		fmt.Println("D")
	}
}
