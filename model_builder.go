package json_schema_tools

import (
	"encoding/json"
	"fmt"
	"go/format"
	"io/ioutil"
//	"log"
)

type Builder struct {
	InputFile  string
	OutputName string
	PkgName    string
	ModelName  string
	SchemaRaw  []byte
	SchemaJSON  map[string]interface{}
	Errors     []string
}

func NewBuilder(inputFile string, modelName string, pkgName string) (builder Builder) {
	builder = Builder{}
	// try to read the given filename
	raw, err := ioutil.ReadFile(inputFile)
	if err != nil {
		msg :=  fmt.Sprintf("reading input file: %s", err)
		builder.Errors = append(builder.Errors, msg)
		return
	}

	builder.InputFile = inputFile
	builder.ModelName = modelName
	builder.PkgName = pkgName
	builder.SchemaRaw = raw

	if err := json.Unmarshal(builder.SchemaRaw, &builder.SchemaJSON); err != nil {
		msg :=  fmt.Sprintf("JSON error: %s", err)
		builder.Errors = append(builder.Errors, msg)
	}
	return
}

func (b *Builder) Generate() ([]byte, error) {
	//TODO add builder.Error without b.SchemaJSON
	src := b.addHeadInfo()

	// add comment from schema[description]
	if b.SchemaJSON["description"] !=nil {
		src += fmt.Sprintf("// %s\n", b.SchemaJSON)
	}

	//TODO if b.PkgName empty set from schmema[title]
	src += fmt.Sprintf("package %s\n", b.PkgName)
	props := NewProperties(b.SchemaJSON)
	//TODO if b.ModelName empty set from schmema[title]
	src += fmt.Sprintf("type %s struct {%s}", b.ModelName, props.ToStruct())
	// TODO add getter & setter methods
	formatted, err := format.Source([]byte(src))
	if err != nil {
		err = fmt.Errorf("error formatting: %s, was formatting\n%s", err, src)
	}
	return formatted, err
}

func (b *Builder) addHeadInfo() string {
	return "//auto-created by generator all changes will be lost.\n" +
		"//\tgo run generator.go -name=Contact -input=contact.json -o=contact.go -pkg=contacts\n\n"
}
