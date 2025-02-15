package main 
import "fmt"

func commandExplore(cfg *config,args ...string) error  {
	if *cfg.areaName=="" {
        return fmt.Errorf("Lütfen bir lokasyon belirtin")
    }
	fmt.Printf("Lokasyon: %s\n",*cfg.areaName)
	exploreResp,err:=cfg.pokeapiClient.ExploreLocation(*cfg.areaName)
	if err!=nil {
        return err
    }
	fmt.Printf("Bulunan pokemonlar %v ve %v",len(exploreResp.Pokemons),exploreResp.Pokemons)
	for _, v := range exploreResp.Pokemons {
		fmt.Println(v.Pokemon.Name)
	}
	return nil
}