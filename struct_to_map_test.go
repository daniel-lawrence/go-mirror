package mirror_test

import (
	"testing"

	mirror "github.com/daniel-lawrence/go-mirror"
)

type testStruct struct {
	unexportedField int
	ExportedInt     int
	ExportedString  string
}

func TestStructToMap(t *testing.T) {
	testObj := testStruct{
		unexportedField: 42,
		ExportedInt:     84,
		ExportedString:  "hello",
	}

	result := mirror.StructToMap(testObj)
	t.Log(result)

	expected := map[string]interface{}{
		"ExportedInt":    testObj.ExportedInt,
		"ExportedString": testObj.ExportedString,
	}
	for k, v := range expected {
		if result[k] != v {
			t.Errorf("%s: was %v, should be %v", k, result[k], v)
		}
	}

	if badConversion := mirror.StructToMap("some string"); badConversion != nil {
		t.Errorf("Expected nil, got %v", badConversion)
	}
}
