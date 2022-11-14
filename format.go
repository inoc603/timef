package timef

import "time"

type Layout struct{}

func (f Layout) String() string { return time.Layout }

type ANSIC struct{}

func (f ANSIC) String() string { return time.ANSIC }

type UnixDate struct{}

func (f UnixDate) String() string { return time.UnixDate }

type RubyDate struct{}

func (f RubyDate) String() string { return time.RubyDate }

type RFC822 struct{}

func (f RFC822) String() string { return time.RFC822 }

type RFC822Z struct{}

func (f RFC822Z) String() string { return time.RFC822Z }

type RFC850 struct{}

func (f RFC850) String() string { return time.RFC850 }

type RFC1123 struct{}

func (f RFC1123) String() string { return time.RFC1123 }

type RFC1123Z struct{}

func (f RFC1123Z) String() string { return time.RFC1123Z }

type RFC3339 struct{}

func (f RFC3339) String() string { return time.RFC3339 }

type RFC3339Nano struct{}

func (f RFC3339Nano) String() string { return time.RFC3339Nano }

type Kitchen struct{}

func (f Kitchen) String() string { return time.Kitchen }

type Stamp struct{}

func (f Stamp) String() string { return time.Stamp }

type StampMilli struct{}

func (f StampMilli) String() string { return time.StampMilli }

type StampMicro struct{}

func (f StampMicro) String() string { return time.StampMicro }

type StampNano struct{}

func (f StampNano) String() string { return time.StampNano }
