// Что выведет код?

package main

import (
	"fmt"
)

const X = 2

func main() {
	const (
		X = X + X
		Y
		Z
		F
	)

	fmt.Println(X, Y, Z, F)
}
