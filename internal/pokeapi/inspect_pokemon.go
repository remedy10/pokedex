package pokeapi


import (
    "encoding/json"
    "io/ioutil"
    "net/http")


func (c *Client)InspectPokemon(pokemonName string) (Pokemon, error) {
	url:=baseURL+"/pokemon/"+pokemonName
	req,err:=http.NewRequest("GET",url,nil)
	if err !=nil{
        return Pokemon{},nil
    }
	resp,err:= c.httpClient.Do(req)
	if err !=nil{
        return Pokemon{},nil
    }
	defer resp.Body.Close()
	data,err := ioutil.ReadAll(resp.Body)
	if err !=nil{
        return Pokemon{},err
    }
	var pokemon Pokemon
	err= json.Unmarshal(data,&pokemon)
	if err !=nil{
        return Pokemon{},nil
    }
	return pokemon, nil
}