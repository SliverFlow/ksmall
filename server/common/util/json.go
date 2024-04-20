package util

import "encoding/json"

// JSONToStruct converts a JSON string to a struct
func JSONToStruct(jsonStr string, v interface{}) error {
	return json.Unmarshal([]byte(jsonStr), v)
}

// StructToJSON converts a struct to a JSON string
func StructToJSON(v interface{}) (string, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(b), err
}
