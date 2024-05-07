package jsonconverter

import "encoding/json"

func FromJSON[T any](bytes []byte) (T, error) {
	var data T

	err := json.Unmarshal(bytes, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
