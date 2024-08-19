# csvreader
CSV Simple Reader

## Requirements
* The first line must be a header
* Al last two columns
* At last three rows (required to determinate the delimiter)

## Features
* Read CSV/TSV files automatically (support delimiter ",", ";", "\t")
* Don't require an deterministic struct to read

## Returned struct
- Headers []string
- Rows    map[int]map[string]string

## Installation
`go get github.com/jonathanhecl/csvreader`

## Example
```go
import (
    "github.com/jonathanhecl/csvreader"
)

func main() {
    data := csvreader.LoadFileCSV("example.csv")
    // ...
    data2 := csvreader.ReadCSV("ID,Name,Extra\n1,Test,2\n2,Another,more\n3,Newest,\n4,Another 2,more data")
    // ...
}

```