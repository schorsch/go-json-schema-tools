package json_schema_tools

import (
	"bitbucket.org/pkg/inflect"
)

// single object property
type Property struct {
	FieldName string
	ValueType string
	Name      string
	RawValues  map[string]interface{}
}

//NewProperty creates a Property from a json string
func NewProperty(name string, values map[string]interface{}) Property {

	prop := Property{
		Name:      name,
		RawValues:  values,
		FieldName: inflect.Camelize(name),
	}
	prop.setType()
	return prop
}


//setType from the type declared in raw values
func (p *Property) setType() {
	switch p.RawValues["type"] {
	case "string":
		p.ValueType = "string"
	case "boolean":
		p.ValueType = "bool"
	case "integer":
		p.ValueType = "int"
	case "number":
		p.ValueType = "float64"
	case "date":
		p.ValueType = "time.Time"
	case "date-time":
		p.ValueType = "time.Time"
	case "array":
		p.ValueType = "[]interface{}"
	case "object":
		p.ValueType = "interface{}"
	default:
		p.ValueType = "string"
	}
}
