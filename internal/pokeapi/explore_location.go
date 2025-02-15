package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)


func (c *Client)ExploreLocation(locationName string) (RespLocation, error)  {
	url:= baseURL+"/location-area/"+locationName
	req,err:= http.NewRequest("GET",url,nil)
	if err!=nil{
        return RespLocation{},nil
    }
	resp,err:= c.httpClient.Do(req)
	if err!=nil{
        return RespLocation{},nil
    }
	defer resp.Body.Close()
	data,err := io.ReadAll(resp.Body)
	if err!=nil{
        return RespLocation{},nil
    }
	locationPokemons:= RespLocation{}
	err= json.Unmarshal(data,&locationPokemons)
	if err!=nil{
        return RespLocation{},nil
    }
	return locationPokemons, nil
}