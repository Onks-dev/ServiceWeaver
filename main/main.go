package main

import (
	context2 "context"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"log"
	"net/http"
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

	lOpts := weaver.ListenerOptions{LocalAddress: "localhost:1234"}
	l, err := root.Listener("reverseListener", lOpts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("reverseListener on %v\n", l)

	http.HandleFunc("/reverse", func(writer http.ResponseWriter, request *http.Request) {
		reversed, err := reverser.Reverse(context, request.URL.Query().Get("name"))
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(writer, "Reversed name, %s\n", reversed)
	})
	http.Serve(l, nil)
}
