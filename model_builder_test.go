package json_schema_tools

import (
//	"encoding/json"
	"testing"
)

//// assert helper
//func assertEqual(t *testing.T, a, b string) {
//	if a != b {
//		t.Errorf("expected %v got %v", a, b)
//	}
//}

func TestNewBuilder(t *testing.T) {

	builder := NewBuilder("test_schema/contact.json", "contact", "contacts")
	assertEqual(t, "contacts", builder.PkgName)
	assertEqual(t, "contact", builder.ModelName)
	assertEqual(t, "test_schema/contact.json", builder.InputFile)
}

func TestNewBuilderError(t *testing.T) {

	builder := NewBuilder("wrong.json", "contact", "contacts")
	assertEqual(t, "reading input file: open wrong.json: no such file or directory", builder.Errors[0])
}