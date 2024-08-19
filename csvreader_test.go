package csvreader

import (
	"reflect"
	"testing"
)

func TestLoadFileCSV(t *testing.T) {
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
			"Load file CSV",
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
			"Load file TSV",
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
			got, err := LoadFileCSV(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadFileCSV() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got.Headers, tt.want.Headers) {
				t.Errorf("LoadFileCSV() Headers = %v, want %v", got.Headers, tt.want.Headers)
			}

			for id, row := range got.Rows {
				if !reflect.DeepEqual(row, tt.want.Rows[id]) {
					t.Errorf("LoadFileCSV() Rows [%d] = %v, want %v", id, row, tt.want.Rows[id])
				}
			}

		})
	}
}

func TestReadCSV(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name    string
		args    args
		want    CSVStruct
		wantErr bool
	}{
		{
			"Read CSV",
			args{content: "ID,Name,Extra\n1,Test,2\n2,Another,more\n3,Newest,\n4,Another 2,more data"},
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
			got, err := ReadCSV(tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadCSV() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got.Headers, tt.want.Headers) {
				t.Errorf("ReadCSV() Headers = %v, want %v", got.Headers, tt.want.Headers)
			}

			for id, row := range got.Rows {
				if !reflect.DeepEqual(row, tt.want.Rows[id]) {
					t.Errorf("ReadCSV() Rows [%d] = %v, want %v", id, row, tt.want.Rows[id])
				}
			}
		})
	}
}
