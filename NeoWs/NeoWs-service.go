package NeoWs

import (
	"nasa-neows-cli-tool/util"
	"time"
)

func ShowNEOsByLastSevenDays() string {
	currentTime := time.Now()
	currentDate := currentTime.Format("2006-01-02")

	sevenDaysAgo := currentTime.AddDate(0, 0, -6)
	sevenDaysAgoDate := sevenDaysAgo.Format("2006-01-02")

	neoWsData, err := getNeoWsByTimePeriod(sevenDaysAgoDate, currentDate)
	util.CheckError(err)

	neoWsFormated := formatNearWsResponse(neoWsData)

	neoWsJson, err := util.ConvertToJSON(neoWsFormated)
	util.CheckError(err)

	return neoWsJson
}

func formatNearWsResponse(nearWsData NeoWsResponse) NeoWs {
	neoWsFormated := NeoWs{
		Total:            nearWsData.ElementCount,
		NearEarthObjects: formatNearEarthObjects(nearWsData.NearEarthObjects),
	}

	return neoWsFormated
}

func formatNearEarthObjects(neoObjects map[string][]NearEarthObject) []NearEarthObjects {
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
