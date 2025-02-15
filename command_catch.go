package main

import (
	"fmt"
	"math/rand/v2"
	"errors"
)
func commandCatch(cfg *config,args ...string) error  {
	fmt.Printf("Throwing a Pokeball at %s...\n",*cfg.areaName)
	pokemon,err:=cfg.pokeapiClient.CatchPokemon(*cfg.areaName)
	if err != nil {
		return err
	}
	if catchLogic(&pokemon.BaseExperience) {
		if _, exists := cfg.pokedex[pokemon.Name];!exists {
			fmt.Printf("%s was caught!\n You may now inspect it with the inspect command.\n",pokemon.Name)
            cfg.pokedex[pokemon.Name]=pokemon
			return nil
        }
		return errors.New(pokemon.Name+" is already caught!\n")
	}
	fmt.Printf("%s escaped!\n",pokemon.Name)
	return nil
}

func catchLogic(experience *int) bool {
	randomNumber:=rand.IntN(500)
	fmt.Printf("Your random number is %d and base_experience is %d\n",randomNumber,*experience)
	if randomNumber < *experience {
        return true
    }
	return false
}