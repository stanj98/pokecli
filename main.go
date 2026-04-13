package main

import "github.com/stanj98/pokecli/internal/pokeapi"

type config struct {
	pokeapiClient          pokeapi.Client
	nextLocationAreaURL    *string
	previousLocationAeaURL *string
}

func main() {

	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
	}

	startRepl(&cfg)
}
