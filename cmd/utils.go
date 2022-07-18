package cmd

import (
	"encoding/json"
)

func JSONify(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}
