package timef_test

import (
	"fmt"
	"time"

	"github.com/inoc603/timef"
)

type CustomFormat struct{}

func (f CustomFormat) String() string {
	return "my custom format: 2006/01/02 15:04"
}

func ExampleTime_customFormat() {
	t, _ := time.Parse(time.RFC1123, "Thu, 18 Aug 2022 19:09:01 CST")

	myTime := timef.New[CustomFormat](t)

	fmt.Println(myTime.String())

	// output: my custom format: 2022/08/18 19:09
}
