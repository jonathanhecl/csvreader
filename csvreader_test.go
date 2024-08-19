package csvreader

import (
	"reflect"
	"testing"
)

func TestCSVReader(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    CSVStruct
		wantErr bool
	}{
		{
			"Load CSV",
			args{filename: "example.csv"},
			CSVStruct{Headers: []string{"ID", "Name", "Extra"},
				Rows: map[int]map[string]string{
					0: {"ID": "1", "Name": "Test", "Extra": "2"},
					1: {"ID": "2", "Name": "Another", "Extra": "more"},
					2: {"ID": "3", "Name": "Newest", "Extra": ""},
					3: {"ID": "4", "Name": "Another 2", "Extra": "more data"},
				}},
			false,
		},
		{
			"Load TSV",
			args{filename: "example.tsv"},
			CSVStruct{Headers: []string{"ID", "Name", "Extra"},
				Rows: map[int]map[string]string{
					0: {"ID": "1", "Name": "Test", "Extra": "2"},
					1: {"ID": "2", "Name": "Another", "Extra": "more"},
					2: {"ID": "3", "Name": "Newest", "Extra": ""},
					3: {"ID": "4", "Name": "Another 2", "Extra": "more data"},
				}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CSVReader(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("CSVReader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got.Headers, tt.want.Headers) {
				t.Errorf("CSVReader() Headers = %v, want %v", got.Headers, tt.want.Headers)
			}

			for id, row := range got.Rows {
				if !reflect.DeepEqual(row, tt.want.Rows[id]) {
					t.Errorf("CSVReader() Rows [%d] = %v, want %v", id, row, tt.want.Rows[id])
				}
			}

		})
	}
}
