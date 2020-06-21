package utils

import "encoding/json"

func ConverStructToString(data interface{}) (string, error) {
	result, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(result), nil
}
