package timef

import (
	"encoding/json"
	"math"
	"time"
)

type Number interface {
	~int | ~int64 | ~int32 | ~uint | ~uint64 | ~uint32
}

// Timestamp is a numberic representation of a time value.
type Timestamp[N Number, U TimestampUnit] struct {
	time.Time
}

func (t *Timestamp[N, F]) UnmarshalJSON(b []byte) error {
	var v N
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}

	var unit F

	t.Time = time.Unix(0, int64(N(v)*N(unit.Unit())))

	return nil
}

func (t Timestamp[N, F]) MarshalJSON() ([]byte, error) {
	var unit F
	v := math.Ceil(float64(t.UnixNano()) / float64(unit.Unit()))
	return json.Marshal(N(v))
}

type TimestampUnit interface {
	Unit() time.Duration
}

type TimestampSeconds struct{}

func (t TimestampSeconds) Unit() time.Duration { return time.Second }

type TimestampMilliseconds struct{}

func (t TimestampMilliseconds) Unit() time.Duration { return time.Millisecond }
