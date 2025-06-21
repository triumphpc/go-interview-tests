package main

import (
	"log"
	"slices"
)

func main() {
	a := []int{1, 2, 3, 2}

	for i, v := range a {
		if v == 2 {
			a = slices.Delete(a, i, i+1)

			continue
		}
		log.Println(v, " ")
	}

}
