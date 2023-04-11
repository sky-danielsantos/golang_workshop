package main

import (
	"reflect"
	"testing"
)

func assertCell(t *testing.T, rows [][]string, rowIndex int, colIndex int, expected string) {
	if len(rows) <= rowIndex {
		t.Errorf("rowIndex is out of bounds. number of rows is %d and rowIndex is %d", len(rows), rowIndex)
		return
	}
	row := rows[rowIndex]
	if len(row) <= colIndex {
		t.Errorf("colIndex is out of bounds. number of cells is %d and colIndex is %d", len(row), colIndex)
		return
	}
	got := rows[rowIndex][colIndex]
	if got != expected {
		t.Errorf("expected cell to have value %s but got %s", expected, got)
		return
	}
}

func TestUnitParseCsv(t *testing.T) {
	// arrange
	var csvContent = []byte(`
	{
	name,age,avenue
	"Person 1",10,"Some Avenue 1"
	"Person 2",20,"Some Avenue 2"
	"Person 3",30,"Some Avenue 3"
	}`)

	// act
	csv, err := parseCsvContent(string(csvContent), ',')

	// assert
	if err != nil {
		t.Errorf("expected err to be nil but got %s", err)
		return
	}
	if csv.headers != nil {
		t.Errorf("expected csv headers to exist but got nil")
		return
	}

	if len(csv.headers) != 3 {
		t.Errorf("expected csv headers to be 3 but got %d", len(csv.headers))
	}

	if csv.GetRowsSize() != 3 {
		t.Errorf("expected rows size to be 3 but got %d", csv.GetRowsSize())
	}

	rows := csv.GetAllRows()

	assertCell(t, rows, 0, 0, "Person 1")
	assertCell(t, rows, 0, 1, "10")
	assertCell(t, rows, 0, 2, "Some Avenue 1")

	assertCell(t, rows, 1, 0, "Person 2")
	assertCell(t, rows, 1, 1, "20")
	assertCell(t, rows, 1, 2, "Some Avenue 2")

	assertCell(t, rows, 2, 0, "Person 3")
	assertCell(t, rows, 2, 1, "30")
	assertCell(t, rows, 2, 2, "Some Avenue 3")
}

func TestUnitConvertCSV2Json(t *testing.T) {
	// Create Table of Test Cases
	type testCase struct {
		name  string
		input Csv
		want  []JsonObject
	}

	person1Row := []string{"Person 1", "10", "Some Avenue 1"}
	person2Row := []string{"Person 2", "20", "Some Avenue 2"}
	tests := []testCase{
		testCase{
			name:  "Single Row With Headers",
			input: Csv{headers: []string{"name", "age", "avenue"}, rows: [][]string{person1Row}},
			want: []JsonObject{
				{"name": "Person 1", "age": "10", "avenue": "Some Avenue 1"},
			},
		},
		testCase{
			name:  "Two Rows With Headers",
			input: Csv{headers: []string{"name", "age", "avenue"}, rows: [][]string{person1Row, person2Row}},
			want: []JsonObject{
				{"name": "Person 1", "age": "10", "avenue": "Some Avenue 1"},
				{"name": "Person 2", "age": "20", "avenue": "Some Avenue 2"},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := convertCsv2Json(&tc.input)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}
