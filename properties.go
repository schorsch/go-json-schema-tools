package json_schema_tools

import (
	"fmt"
	"sort"
)

// Properties is a slice with all properties of an object, used for sorting properties by name
type Properties []*Property

// NewProperties builds a properties map for the given raw json object
func NewProperties(obj map[string]interface{}) Properties {
	props, ok := obj["properties"]
	if !ok {
		// wrong markup or no props
		// TODO log error ? is it ok to return an empty slice?
		return make(Properties, 0, 0)
	}
	fields := props.(map[string]interface{})
	if fields == nil {
		// no props
		return make(Properties, 0, 0)
	}
	properties := make(Properties, 0, len(fields))
	// for each property
	for name, values := range fields {
        prop := NewProperty(name, values.(map[string]interface{}) )
		properties = append(properties, &prop)
	}
	return properties
}

func (ps *Properties) ToStruct() string {
	// sort the properties
	sort.Sort(ps)
	output := ""
	for _, p := range *ps {
		output += fmt.Sprintf("\n%s %s `json:\"%s\"`",
			p.FieldName,
			p.ValueType,
			p.Name)
	}
	return output
}

// ======== Sorting
// Len is part of sort.Interface.
func (p Properties) Len() int {
	return len(p)
}

// Swap is part of sort.Interface.
func (p Properties) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Less is part of sort.Interface. We use FieldName as the value to sort by
func (p Properties) Less(i, j int) bool {
	return p[i].FieldName < p[j].FieldName
}