package util

import (
	"encoding/json"
)

func ConvertFromJson[T any](bytes []byte) T {
	var data T
	if err := json.Unmarshal(bytes, &data); err != nil {
		panic(err)
	}
	return data
}
