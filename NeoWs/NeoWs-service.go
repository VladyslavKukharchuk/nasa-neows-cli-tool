package NeoWs

import (
	"nasa-neows-cli-tool/util"
	"time"
)

func GetNEOsByDaysAgo(daysCount int) string {
	neoWsResponsesCh := make(chan NeoWsResponse, daysCount)

	go func() {
		currentTime := time.Now()

		for dayNumber := 0; dayNumber < daysCount; dayNumber++ {
			selectedDay := currentTime.AddDate(0, 0, -dayNumber)
			selectedDayDate := selectedDay.Format("2006-01-02")

			neoWsData, err := getNeoWsByTimePeriod(selectedDayDate, selectedDayDate)
			util.CheckError(err)

			neoWsResponsesCh <- neoWsData
		}

		close(neoWsResponsesCh)
	}()

	var neoWsResponses []NeoWsResponse
	for neoWsResponse := range neoWsResponsesCh {
		neoWsResponses = append(neoWsResponses, neoWsResponse)
	}

	neoWsFormated := formatNearWsResponses(neoWsResponses)
	neoWsJson, err := util.ConvertToJSON(neoWsFormated)
	util.CheckError(err)

	return neoWsJson
}

func formatNearWsResponses(neoWsData []NeoWsResponse) NeoWs {
	var total int
	var neoObjectsFormated []NearEarthObjects

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
	var neoObjectsFormated []NearEarthObjects
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
