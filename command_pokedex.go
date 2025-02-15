package main 

import("fmt")

func commandPokedex(c *config,arg ...string) error{
	if len(c.pokedex)==0 {
		fmt.Printf("Your pokedex is empty!\n")
	}
	fmt.Printf("Your Pokedex:\n")
	for k, _ := range c.pokedex {
		fmt.Printf("\t-%s\n",k)
	}
	return nil
}