package neows

import (
	"encoding/json"
	"time"
)

func GetNEOsByDaysAgo(URL string, apiKey string, count int) (string, error) {
	dates := getDates(count)

	neoWs := GetNEOsByDates(URL, apiKey, dates)
	neoWsJSON, err := json.Marshal(neoWs)
	if err != nil {
		return "", err
	}

	return string(neoWsJSON), nil
}

func getDates(days int) []string {
	dates := make([]string, 0)
	currentTime := time.Now()

	for i := 0; i < days; i++ {
		date := currentTime.AddDate(0, 0, -i)
		dateStr := date.Format("2006-01-02")
		dates = append(dates, dateStr)
	}

	return dates
}

func GetNEOsByDates(URL string, apiKey string, dates []string) NeoWs {
	neoWsResponsesCh := make(chan *NeoWsResponse)

	go func() {
		for _, date := range dates {
			neoWsData, err := getNeoWsByTimePeriod(URL, date, date, apiKey)
			if err != nil {
				panic(err)
			}

			neoWsResponsesCh <- neoWsData
		}

		close(neoWsResponsesCh)
	}()

	neoWsResponses := make([]*NeoWsResponse, 0)

	for neoWsResponse := range neoWsResponsesCh {
		neoWsResponses = append(neoWsResponses, neoWsResponse)
	}

	neoWs := FormatNearWsResponses(neoWsResponses)

	return neoWs
}

func FormatNearWsResponses(neoWsData []*NeoWsResponse) NeoWs {
	var total int
	neoObjectsFormated := make([]NearEarthObjects, 0)

	for _, data := range neoWsData {
		total += data.ElementCount
		neoObjectsFormated = append(neoObjectsFormated, formatNearObjects(data.NearEarthObjects)...)
	}

	return NeoWs{
		Total:            total,
		NearEarthObjects: neoObjectsFormated,
	}
}

func formatNearObjects(neoObjects map[string][]NearEarthObject) []NearEarthObjects {
	neoObjectsFormated := make([]NearEarthObjects, 0)

	for date, neoObjectsByDay := range neoObjects {
		for _, neoObject := range neoObjectsByDay {
			neoObj := NearEarthObjects{
				Date:                           date,
				ID:                             neoObject.ID,
				Name:                           neoObject.Name,
				IsPotentiallyHazardousAsteroid: neoObject.IsPotentiallyHazardousAsteroid,
			}
			neoObjectsFormated = append(neoObjectsFormated, neoObj)
		}
	}

	return neoObjectsFormated
}
