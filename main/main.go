package main

import (
	context2 "context"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"log"
)

func main() {
	// Initialize service weaver
	context := context2.Background()
	root := weaver.Init(context)

	//Get a client to the reverser component
	reverser, err := weaver.Get[Reverser](root)
	if err != nil {
		log.Fatal(err)
	}

	reversed, err := reverser.Reverse(context, "Khumo")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reversed)
}
