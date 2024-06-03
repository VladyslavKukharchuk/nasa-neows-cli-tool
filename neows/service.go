package neows

import (
	"encoding/json"
	"golang.org/x/sync/errgroup"
	"sync"
	"time"
)

func GetNEOsByDaysAgo(URL string, apiKey string, count int) (string, error) {
	dates := getDates(count)

	neoWs, err := GetNEOsByDates(URL, apiKey, dates)
	if err != nil {
		return "", err
	}

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

func GetNEOsByDates(URL string, apiKey string, dates []string) (NeoWs, error) {
	var neoWsResponses = make([]*NeoWsResponse, 0, len(dates))
	var mu sync.Mutex
	var g = errgroup.Group{}

	for _, date := range dates {
		date := date

		g.Go(func() error {
			neoWsData, err := getNeoWsByTimePeriod(URL, date, date, apiKey)
			if err != nil {
				return err
			}

			mu.Lock()
			neoWsResponses = append(neoWsResponses, neoWsData)
			mu.Unlock()

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return NeoWs{}, err
	}

	neoWs := FormatNearWsResponses(neoWsResponses)

	return neoWs, nil
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
