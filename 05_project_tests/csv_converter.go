package main

import (
	"encoding/json"
	"os"
	"strings"
)

// New fields can be added
// New functions can be added
type Csv struct {
	headers []string
	rows    [][]string
}

func (c *Csv) GetAllRows() [][]string {
	return c.rows
}

func (c *Csv) GetRow(index int) []string {
	return c.rows[index]
}

func (c *Csv) GetHeaders() []string {
	return c.headers
}

func (c *Csv) GetRowsSize() int {
	return len(c.rows)
}

type JsonObject map[string]string

func (j JsonObject) AddKeyValue(key string, value string) {
	j[key] = value
}

func NewJsonObject() JsonObject {
	return make(JsonObject)
}

func splitString(s string, delem string) []string {
	return strings.Split(s, delem)
}

func trimQuote(s string) string {
	if len(s) > 0 && s[0] == '"' {
		s = s[1:]
	}
	if len(s) > 0 && s[len(s)-1] == '"' {
		s = s[:len(s)-1]
	}
	return s
}

// TODO: implement me ( assume that csv has always at least one row and one header )
func parseCsvContent(csvContent string, delim string) Csv {
	return Csv{}
}

// TODO: implement me
func convertCsv2Json(csv Csv) []JsonObject {
	var listOfJsonObjects []JsonObject
	obj := JsonObject{"k1": "v"}
	obj.AddKeyValue("k2", "v")
	listOfJsonObjects = append(listOfJsonObjects, obj)
	return listOfJsonObjects
}

func ReadAndConvert2Json(filepath string, delim string) (string, error) {
	csvContent, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return ConvertCsv2Json(string(csvContent), delim)
}

// TODO: implement me (needs to be converted/marshalled to JSON)
func ConvertCsv2Json(csvContent string, delim string) (string, error) {
	csv := parseCsvContent(csvContent, delim)
	jsonObjectList := convertCsv2Json(csv)
	json, err := json.MarshalIndent(jsonObjectList, "", "  ")
	if err != nil {
		return "", err
	}
	return string(json), nil
}
