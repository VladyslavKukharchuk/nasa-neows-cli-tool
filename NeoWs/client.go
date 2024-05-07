package NeoWs

import (
	"errors"
	"io"
	"nasa-neows-cli-tool/util"
	"net/http"
	"net/url"
)

func getNeoWsByTimePeriod(URL, startDate, endDate, apiKey string) (NeoWsResponse, error) {
	baseURL, err := url.Parse(URL + "feed")
	if err != nil {
		return NeoWsResponse{}, err
	}

	params := url.Values{}
	params.Add("start_date", startDate)
	params.Add("end_date", endDate)
	params.Add("api_key", apiKey)
	baseURL.RawQuery = params.Encode()

	resp, err := http.Get(baseURL.String())
	if err != nil {
		return NeoWsResponse{}, err
	}

	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return NeoWsResponse{}, err
	}

	if resp.StatusCode != http.StatusOK {
		apiError := util.ConvertFromJSON[NeoWsError](bytes)
		return NeoWsResponse{}, errors.New(apiError.Error.Message)
	}

	neoWsData := util.ConvertFromJSON[NeoWsResponse](bytes)

	return neoWsData, nil
}
