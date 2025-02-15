package main

import (
	"fmt"
	"strconv"
	)
	

func commandInspect(cfg *config, args ...string) error {
	pokemon, err := cfg.pokeapiClient.InspectPokemon(*cfg.areaName)
	if err != nil {
		return err
	}
	message := fmt.Sprintf("Name: %s\nHeight: %d\nWeight: %d\nStats:", pokemon.Name, pokemon.Height, pokemon.Weight)
	for _, v := range pokemon.Stats {
		message += "\n\t-" + v.Stat.Name + ": " + strconv.Itoa(v.BaseStat)
	}
	message += "\nTypes:"
	for _, v := range pokemon.Types {
		message += "\n\t- " + v.Type.Name
	}
	fmt.Println(message)
	return nil
}

