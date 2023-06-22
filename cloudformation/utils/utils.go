package utils

import (
	"encoding/json"
	"strings"
)

// ByJSONLength implements the sort interface and enables sorting by JSON length
type ByJSONLength []interface{}

func (s ByJSONLength) Len() int {
	return len(s)
}

func (s ByJSONLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByJSONLength) Less(i, j int) bool {
	// Nil is always at the end
	if s[i] == nil {
		return false
	}
	if s[j] == nil {
		return true
	}
	jsoni, _ := json.Marshal(s[i])
	jsonj, _ := json.Marshal(s[j])

	if jsoni == nil {
		return false
	}
	if jsonj == nil {
		return true
	}

	return len(jsoni) > len(jsonj)
}

// GetStringInBetweenTwoString is a function that will grab the string value between two provided strings
func GetStringInBetweenTwoString(str string, startS string, endS string) (result string, found bool) {
	s := strings.Index(str, startS)
	if s == -1 {
		return result, false
	}
	newS := str[s+len(startS):]
	e := strings.Index(newS, endS)
	if e == -1 {
		return result, false
	}
	result = newS[:e]
	return result, true
}
