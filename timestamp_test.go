package timef

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func testFormat[F TimestampUnit](r *require.Assertions) {
	var format F
	now := time.Now().Truncate(format.Unit())

	nowBytes, err := json.Marshal(int64(float64(now.UnixNano()) / float64(format.Unit())))
	r.Nil(err)

	var ts Timestamp[int, F]
	r.Nil(json.Unmarshal(nowBytes, &ts))

	r.Equal(now.UnixNano(), ts.UnixNano())

	tsBytes, err := json.Marshal(ts)
	r.Nil(err)

	r.Equal(nowBytes, tsBytes)

	r.NotNil(json.Unmarshal([]byte(`false`), &ts))
}

func TestTimestamp(t *testing.T) {
	r := require.New(t)

	testFormat[TimestampSeconds](r)
	testFormat[TimestampMilliseconds](r)
}
