package pokeapi

// RespShallowLocations -
type RespShallowLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type RespLocation struct {
	Pokemons []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
		}

type Pokemon struct {
	Name string `json:"name"`
	BaseExperience int `json:"base_experience"`
	Height int `json:"height"`
	Weight int `json:"weight"`
	Stats []Stat `json:"stats"`
	Types []Types `json:"types"`
}

type NameFor struct {
	Name string `json:"name"`
}

type Stat struct {
	BaseStat int `json:"base_stat"`
	Stat NameFor `json:"stat"`
}
type Types struct {
	Type NameFor `json:"type"`
}