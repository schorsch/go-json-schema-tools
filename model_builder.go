package json_schema_tools

import (
	"bitbucket.org/pkg/inflect"
	"encoding/json"
	"fmt"
	"go/format"
	"io/ioutil"
	"strings"
)

type Builder struct {
	InputFile  string
	OutputName string
	PkgName    string
	ModelName  string
	SchemaRaw  []byte
	SchemaJSON map[string]interface{}
	Errors     []string
}

//NewBuilder creates a build instance using the JSON schema data from the given filename. If modelName and/or pkgName
//are left empty the title-attribute of the JSON schema is used for:
// => struct-name (singular, camelcase e.g contact => Contact)
// => package name (pluralize, lowercased e.g. payment_reminder => paymentreminders)
func NewBuilder(inputFile string, modelName string, pkgName string) (builder Builder) {
	builder = Builder{}
	// try to read input file
	raw, err := ioutil.ReadFile(inputFile)
	if err != nil {
		msg := fmt.Sprintf("File error: %s", err)
		builder.Errors = append(builder.Errors, msg)
		return
	}
	builder.InputFile = inputFile
	builder.SchemaRaw = raw
	// try parsing json
	if err := json.Unmarshal(builder.SchemaRaw, &builder.SchemaJSON); err != nil {
		msg := fmt.Sprintf("JSON error: %s", err)
		builder.Errors = append(builder.Errors, msg)
		return
	}

	// defer model name from schema.title if not given as argument, schema['title'] MUST be set
	if len(modelName) > 0 {
		builder.ModelName = modelName
	} else {
		builder.ModelName = inflect.Typeify(builder.SchemaJSON["title"].(string))
	}
	// defer package name from schema.title if not given as argument
	if len(pkgName) > 0 {
		builder.PkgName = pkgName
	} else {
		//Pluralize no underscores
		builder.PkgName = strings.ToLower(inflect.Camelize(inflect.Pluralize(builder.SchemaJSON["title"].(string))))
	}

	return
}

// Generate creates the golang source file with the struct(model) information from a schema
func (b *Builder) Generate() (formatted []byte) {

	if len(b.SchemaJSON) == 0 {
		msg := fmt.Sprintf("Cannot generate output without a Schema")
		b.Errors = append(b.Errors, msg)
		return formatted
	}

	src := b.addHeadInfo()

	// add comment from schema[description]
	if b.SchemaJSON["description"] != nil {
		src += fmt.Sprintf("// %s\n", b.SchemaJSON["description"])
	}
	src += fmt.Sprintf("package %s\n", b.PkgName)
	props := NewProperties(b.SchemaJSON)
	src += fmt.Sprintf("type %s struct {%s}", b.ModelName, props.ToStruct())
	formatted, err := format.Source([]byte(src))
	if err != nil {
		msg := fmt.Sprintf("error formatting: %s, was formatting\n%s", err, src)
		b.Errors = append(b.Errors, msg)
	}
	return formatted
}

func (b *Builder) addHeadInfo() string {
	// TODO log run command with real used options/flags
	return "//auto-created by generator all changes will be lost.\n" +
		"//\tgo run generator.go -name=Contact -input=contact.json -o=contact.go -pkg=contacts\n\n"
}
