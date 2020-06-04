package main

import (
	"fmt"
	"math/rand"
)

func main() {
	s1 := rand.Perm(12)
	fmt.Println(s1)
}
