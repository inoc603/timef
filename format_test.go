package timef

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestStdFormat(t *testing.T) {
	r := require.New(t)
	r.Equal(Layout{}.String(), time.Layout)
	r.Equal(ANSIC{}.String(), time.ANSIC)
	r.Equal(UnixDate{}.String(), time.UnixDate)
	r.Equal(RubyDate{}.String(), time.RubyDate)
	r.Equal(RFC822{}.String(), time.RFC822)
	r.Equal(RFC822Z{}.String(), time.RFC822Z)
	r.Equal(RFC850{}.String(), time.RFC850)
	r.Equal(RFC1123{}.String(), time.RFC1123)
	r.Equal(RFC1123Z{}.String(), time.RFC1123Z)
	r.Equal(RFC3339{}.String(), time.RFC3339)
	r.Equal(RFC3339Nano{}.String(), time.RFC3339Nano)
	r.Equal(Kitchen{}.String(), time.Kitchen)
	r.Equal(Stamp{}.String(), time.Stamp)
	r.Equal(StampMilli{}.String(), time.StampMilli)
	r.Equal(StampMicro{}.String(), time.StampMicro)
	r.Equal(StampNano{}.String(), time.StampNano)
}
