package util

import "encoding/json"

func ConvertFromJSON[T any](bytes []byte) T {
	var data T
	err := json.Unmarshal(bytes, &data)
	CheckError(err)
	return data
}
