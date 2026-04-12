package main

import (
	"fmt"
	"log"

	"github.com/stanj98/pokecli/internal/pokeapi"
)

func main() {
	pokeAPIClient := pokeapi.NewClient()
	resp, err := pokeAPIClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
	// startRepl()
}
