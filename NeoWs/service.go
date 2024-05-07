package NeoWs

import "nasa-neows-cli-tool/util"

func GetNEOsByDaysAgo(count int) string {
	dates := util.GetDates(count)

	neoWs := GetNEOsByDates(dates)
	neoWsJson := util.ConvertToJSON(neoWs)

	return neoWsJson
}

func GetNEOsByDates(dates []string) NeoWs {
	neoWsResponsesCh := make(chan NeoWsResponse, len(dates))

	go func() {
		for _, date := range dates {
			neoWsData, err := getNeoWsByTimePeriod(date, date)
			util.CheckError(err)

			neoWsResponsesCh <- neoWsData
		}

		close(neoWsResponsesCh)
	}()

	var neoWsResponses []NeoWsResponse
	for neoWsResponse := range neoWsResponsesCh {
		neoWsResponses = append(neoWsResponses, neoWsResponse)
	}

	neoWs := FormatNearWsResponses(neoWsResponses)

	return neoWs
}

func FormatNearWsResponses(neoWsData []NeoWsResponse) NeoWs {
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
