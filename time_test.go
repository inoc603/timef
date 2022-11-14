package timef

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const timeStr = "Thu, 18 Aug 2022 19:09:01 CST"

func TestJson(t *testing.T) {

	jsonStr := `{"time":"` + timeStr + `"}`

	r := require.New(t)

	var payload struct {
		Time Time[RFC1123] `json:"time"`
	}

	r.Nil(json.Unmarshal([]byte(jsonStr), &payload))

	marshaled, err := json.Marshal(payload)
	r.Nil(err)

	r.Equal(jsonStr, string(marshaled))

	r.Equal(timeStr, payload.Time.String())

	buf := bytes.NewBuffer(nil)
	fmt.Fprint(buf, payload.Time)
	r.Equal(timeStr, buf.String())

	r.NotNil(json.Unmarshal([]byte(`{"time":1}`), &payload))

	r.NotNil(json.Unmarshal([]byte(`{"time":"xx"}`), &payload))
}

func TestParse(t *testing.T) {
	r := require.New(t)

	ts, err := Parse[RFC1123](timeStr)
	r.Nil(err)
	r.Equal(timeStr, Format[RFC1123](ts.Time))

	{
		ts, err := Parse[RFC1123]("xx")
		r.Nil(ts)
		r.NotNil(err)
	}
}

func TestNew(t *testing.T) {
	r := require.New(t)

	ts := time.Now()

	formated := New[RFC1123](ts)

	r.Equal(ts, formated.Time)
	r.Equal(ts.Format(time.RFC1123), formated.String())
}
