package utils

import (
	"encoding/json"
)

func ToJson(v interface{}) string {
	if v == nil {
		return ""
	}
	b, _ := json.Marshal(v)
	return string(b)
}
func ToObj(v interface{}, b []byte) {
	json.Unmarshal(b, &v)

}
