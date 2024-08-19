package csvreader

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CSVStruct struct {
	delimiter string
	Headers   []string                  // Headers[0] = "label"
	Rows      map[int]map[string]string // Rows[0]["label"] = "value"
}

var delimiters = []string{",", ";", "\t", ""}

func (data *CSVStruct) addRow(id int, row string) {
	items := strings.Split(row, data.delimiter)
	data.Rows[id] = make(map[string]string)
	for i := 0; i < len(data.Headers); i++ {
		if i < len(items) {
			data.Rows[id][data.Headers[i]] = items[i]
		} else {
			data.Rows[id][data.Headers[i]] = ""
		}
	}
}

func CSVReader(filename string) (CSVStruct, error) {
	csvFile, err := os.Open(filename)
	if err != nil {
		return CSVStruct{}, err
	}
	defer csvFile.Close()

	var data CSVStruct = CSVStruct{
		delimiter: "",
		Headers:   make([]string, 0),
		Rows:      make(map[int]map[string]string, 0),
	}
	var bufferLines []string
	var id int = 0

	reader := bufio.NewScanner(csvFile)
	reader.Split(bufio.ScanLines)

	for reader.Scan() {
		line := reader.Text()

		if line == "" {
			continue
		}

		if len(bufferLines) < 3 {
			bufferLines = append(bufferLines, line)
		}

		if len(bufferLines) == 3 && data.delimiter == "" {
			for _, d := range delimiters {
				s0 := strings.Split(bufferLines[0], d)
				s1 := strings.Split(bufferLines[1], d)
				s2 := strings.Split(bufferLines[2], d)
				if len(s0) > 1 && len(s0) == len(s1) && len(s0) == len(s2) {
					data.delimiter = d
					data.Headers = s0
					break
				}
			}

			if data.delimiter == "" {
				return CSVStruct{}, fmt.Errorf("delimiter not found")
			}

			data.addRow(0, bufferLines[1])
			data.addRow(1, bufferLines[2])
			continue
		} else if data.delimiter != "" {
			data.addRow(id, line)
		}
		id++
	}

	fmt.Println("Delimiter:", data.delimiter)
	fmt.Println("Headers:", data.Headers)
	fmt.Println("Rows:", data.Rows)

	return data, nil
}
