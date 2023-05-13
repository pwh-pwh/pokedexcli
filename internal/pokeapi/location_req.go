package pokeapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *Client) LocationGet(locationName string) (Location, error) {
	url := baseUrl + "/location-area/" + locationName
	var location Location
	if data, exists := c.cache.Get(url); exists {
		err := json.Unmarshal(data, &location)
		if err != nil {
			return location, err
		}
	}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return location, nil
	}
	resp, err := c.httpClient.Do(request)
	if err != nil {
		return location, err
	}
	if resp.StatusCode > 399 {
		return location, fmt.Errorf("resp statuscode error:%v", resp.StatusCode)
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return location, err
	}
	err = json.Unmarshal(bytes, &location)
	if err != nil {
		return location, err
	}
	c.cache.Add(url, bytes)
	return location, nil
}
