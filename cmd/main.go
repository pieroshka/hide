package main

import (
	"fmt"

	"github.com/pieroshka/hide"
)

func main() {
	s := "hello world"

	concealer := hide.New()
	hidden, err := concealer.Hide(s)
	if err != nil {
		panic(err)
	}

	unhidden, err := concealer.Unhide(hidden)
	if err != nil {
		panic(err)
	}

	fmt.Printf(">%s<\n", s)
	fmt.Printf(">%s<\n", hidden)
	fmt.Printf(">%s<\n", unhidden)
}
