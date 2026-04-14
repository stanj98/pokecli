# PokeCLI

A simple and fast command-line interface to explore Pokémon data directly from your terminal.  
Built on top of the PokeAPI, PokeCLI lets you fetch detailed information about Pokémon, abilities, moves, items, and more.

---

## Features

- Search Pokémon, moves, abilities, and items
- View detailed stats, types, and abilities

## Installation

### Using `go install`

```bash
go install github.com/stanj98/pokecli@latest

git clone https://github.com/stanj98/pokecli.git
cd pokecli
go build -o poke-cli
```

## Usage
```bash
poke-cli <command> [flags]
```