package pokeapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *Client) ListLocationAreas(pageUrl *string) (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullUrl := baseUrl + endpoint
	if pageUrl != nil {
		fullUrl = *pageUrl
	}
	if data, exists := c.cache.Get(fullUrl); exists {
		resp := LocationAreasResp{}
		err := json.Unmarshal(data, &resp)
		if err != nil {
			return resp, err
		}
		return resp, nil
	}
	request, err := http.NewRequest(http.MethodGet, fullUrl, nil)
	var locationAreas LocationAreasResp
	if err != nil {
		return locationAreas, err
	}
	response, err := c.httpClient.Do(request)
	if err != nil {
		return locationAreas, err
	}
	defer response.Body.Close()
	if response.StatusCode > 399 {
		return locationAreas, fmt.Errorf("bad status code:%v", response.Status)
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return locationAreas, err
	}
	err = json.Unmarshal(bytes, &locationAreas)
	if err != nil {
		return locationAreas, err
	}
	c.cache.Add(fullUrl, bytes)
	return locationAreas, nil
}
