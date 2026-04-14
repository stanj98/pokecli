package main

import (
	"time"

	"github.com/stanj98/pokecli/internal/pokeapi"
)

type config struct {
	pokeapiClient          pokeapi.Client
	nextLocationAreaURL    *string
	previousLocationAeaURL *string
	caughtPokemon          map[string]pokeapi.Pokemon
}

func main() {

	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}

	startRepl(&cfg)
}
