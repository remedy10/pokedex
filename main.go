package main
import (
	"time"

	"github.com/bootdotdev/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	yourPokedex:=make(map[string]pokeapi.Pokemon)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex:yourPokedex,
	}

	startRepl(cfg)
}
