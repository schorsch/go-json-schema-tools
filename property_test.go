package json_schema_tools

import (
	"encoding/json"
	"testing"
)

// assert helper
func assertEqual(t *testing.T, a, b string) {
	if a != b {
		t.Errorf("expected %v got %v", a, b)
	}
}

func TestNewProperty(t *testing.T) {
	name := "type"
	json_str := []byte(`{"description": "Type of contact",
							"enum":["Client", "Lead", "Supplier"],
							"required" : true,
							"type":"string",
							"maxLength": 50
							}`)
	var raw_prop interface{}
	json.Unmarshal(json_str, &raw_prop)

	p := NewProperty(name, raw_prop.(map[string]interface{}))

	assertEqual(t, "type", p.Name)
	assertEqual(t, "Type", p.FieldName)
	assertEqual(t, "string", p.ValueType)
}
