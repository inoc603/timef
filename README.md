# timef

[![Go Reference](https://pkg.go.dev/badge/github.com/inoc603/timef.svg)](https://pkg.go.dev/github.com/inoc603/timef)

> Convinient time parsing and formatting with generics.

## Usage

### JSON

```go
func ExampleTime_json() {
	var payload struct {
		Time timef.Time[timef.RFC1123] `json:"time"`
	}

	jsonStr := `{"time":"Thu, 18 Aug 2022 19:09:01 CST"}`

	_ = json.Unmarshal([]byte(jsonStr), &payload)

	fmt.Println(payload.Time)

	// output: Thu, 18 Aug 2022 19:09:01 CST
}
```

### Parse

```go
func ExampleParse() {
	t, _ := timef.Parse[timef.RFC1123]("Tue, 15 Nov 2022 11:21:42 CST")

	fmt.Println(t.UTC())

	// output: 2022-11-15 03:21:42 +0000 UTC
}
```

### Custom Format

```go
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
```
