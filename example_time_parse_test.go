package timef_test

import (
	"fmt"

	"github.com/inoc603/timef"
)

func ExampleParse() {
	t, _ := timef.Parse[timef.RFC1123]("Tue, 15 Nov 2022 11:21:42 CST")

	fmt.Println(t.UTC())

	// output: 2022-11-15 03:21:42 +0000 UTC
}
