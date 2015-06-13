package json_schema_tools

import (
	//	"encoding/json"
	"strings"
	"testing"
)

func TestNewBuilderWithNames(t *testing.T) {

	builder := NewBuilder("test_schema/contact.json", "Contact", "contacts")
	assertEqual(t, "contacts", builder.PkgName)
	assertEqual(t, "Contact", builder.ModelName)
	assertEqual(t, "test_schema/contact.json", builder.InputFile)
	if len(builder.SchemaJSON) == 0 {
		t.Error("Expected SchemaJSON, got empty value", builder.SchemaJSON)
	}
}

func TestNewBuilderWithSchemaTitle(t *testing.T) {
	builder := NewBuilder("test_schema/contact.json", "", "")
	assertEqual(t, "contacts", builder.PkgName)
	assertEqual(t, "Contact", builder.ModelName)
}

func TestNewBuilderError(t *testing.T) {
	builder := NewBuilder("wrong.json", "contact", "contacts")
	assertEqual(t, "File error: open wrong.json: no such file or directory", builder.Errors[0])

	builder2 := NewBuilder("test_schema/invalid_json.json", "contact", "contacts")
	if !strings.Contains(builder2.Errors[0], "JSON error:") {
		t.Error("Expected JSON error, got: ", builder2.Errors[0])
	}
}

func TestGenerate(t *testing.T) {

	builder := NewBuilder("test_schema/contact.json", "Contact", "contacts")
	str := builder.Generate()
	if !strings.Contains(string(str), "package contacts") {
		t.Error("Expected package definition, got: ", string(str))
	}
	if !strings.Contains(string(str), "type Contact struct {") {
		t.Error("Expected package definition, got: ", string(str))
	}
}
