package neows

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

type ClientInterface interface {
	GetNeoWsByTimePeriod(startDate, endDate string) (*NeoWsResponse, error)
}

type Client struct {
	apiKey string
}

func NewClient(apiKey string) Client {
	return Client{apiKey: apiKey}
}

func (c *Client) GetNeoWsByTimePeriod(startDate, endDate string) (*NeoWsResponse, error) {
	const URL = "https://api.nasa.gov/neo/rest/v1/"

	baseURL, err := url.Parse(URL + "feed")
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Add("start_date", startDate)
	params.Add("end_date", endDate)
	params.Add("api_key", c.apiKey)
	baseURL.RawQuery = params.Encode()

	resp, err := http.Get(baseURL.String())
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		var apiError NeoWsError

		err := json.Unmarshal(bytes, &apiError)
		if err != nil {
			return nil, err
		}

		return nil, errors.New(apiError.Error.Message)
	}

	var neoWsData NeoWsResponse
	err = json.Unmarshal(bytes, &neoWsData)
	if err != nil {
		return nil, err
	}

	return &neoWsData, nil
}
