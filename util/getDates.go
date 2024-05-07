package util

import "time"

func GetDates(days int) []string {
	dates := make([]string, days)
	currentTime := time.Now()

	for i := 0; i < days; i++ {
		date := currentTime.AddDate(0, 0, -i)
		dateStr := date.Format("2006-01-02")
		dates = append(dates, dateStr)
	}

	return dates
}
