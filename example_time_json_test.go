package timef_test

import (
	"encoding/json"
	"fmt"

	"github.com/inoc603/timef"
)

func ExampleTime_json() {
	var payload struct {
		Time timef.Time[timef.RFC1123] `json:"time"`
	}

	jsonStr := `{"time":"Thu, 18 Aug 2022 19:09:01 CST"}`

	_ = json.Unmarshal([]byte(jsonStr), &payload)

	fmt.Println(payload.Time)

	// output: Thu, 18 Aug 2022 19:09:01 CST
}
