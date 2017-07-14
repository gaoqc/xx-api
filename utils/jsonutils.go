package utils

import (
	"encoding/json"
)

func ToJson(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}
func toObj(v interface{}, jsonStr string) {
	json.Unmarshal([]byte(jsonStr), &v)
}
