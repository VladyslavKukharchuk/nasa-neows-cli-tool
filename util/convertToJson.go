package util

import "encoding/json"

func ConvertToJSON(data interface{}) string {
	jsonData, err := json.Marshal(data)
	CheckError(err)
	return string(jsonData)
}
