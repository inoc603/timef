# timef

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
