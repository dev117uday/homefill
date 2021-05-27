package functions

import "encoding/json"

func ConvertToStruct(data []byte, toStruct interface{}) interface{} {
	json.Unmarshal(data, toStruct)
	return toStruct
}
