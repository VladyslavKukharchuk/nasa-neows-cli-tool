package NeoWs

import (
	"errors"
	"io"
	"nasa-neows-cli-tool/util"
	"net/http"
	"os"
)

const URL = "https://api.nasa.gov/neo/rest/v1/"

var apiKey = os.Getenv("API_KEY")

func getNeoWsByTimePeriod(startDate string, endDate string) (NeoWsResponse, error) {
	resp, err := http.Get(URL + "feed?start_date=" + startDate + "&end_date=" + endDate + "&api_key=" + apiKey)
	util.CheckError(err)

	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	util.CheckError(err)

	if resp.StatusCode != 200 {
		apiError := util.ConvertFromJSON[NeoWsError](bytes)
		return NeoWsResponse{}, errors.New(apiError.Error.Message)
	}

	neoWsData := util.ConvertFromJSON[NeoWsResponse](bytes)

	return neoWsData, nil
}
