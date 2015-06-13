package json_schema_tools

import (
	"encoding/json"
	"fmt"
	"strings"
	"go/format"
	"io/ioutil"
	"bitbucket.org/pkg/inflect"
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

//NewBuilder creates a build instance using the JSON schema data from the given filename. If modelName and/or pkgName
//are left empty the title-attribute of the JSON schema is used for:
// => struct-name (singular, camelcase e.g contact => Contact)
// => package name (pluralize, lowercased e.g. payment_reminder => paymentreminders)
func NewBuilder(inputFile string, modelName string, pkgName string) (builder Builder) {
	builder = Builder{}
	// try to read input file
	raw, err := ioutil.ReadFile(inputFile)
	if err != nil {
		msg :=  fmt.Sprintf("File error: %s", err)
		builder.Errors = append(builder.Errors, msg)
		return
	}
	builder.InputFile = inputFile
	builder.SchemaRaw = raw
	// try parsing json
	if err := json.Unmarshal(builder.SchemaRaw, &builder.SchemaJSON); err != nil {
		msg :=  fmt.Sprintf("JSON error: %s", err)
		builder.Errors = append(builder.Errors, msg)
		return
	}

	// defer model name from schema.title if not given as argument, schema['title'] MUST be set
	if len(modelName) > 0 {
		builder.ModelName = modelName
	}else{
		builder.ModelName = inflect.Typeify(builder.SchemaJSON["title"].(string))
	}
	// defer package name from schema.title if not given as argument
	if len(pkgName) > 0 {
		builder.PkgName = pkgName
	}else{
		//Pluralize no underscores
		builder.PkgName = strings.ToLower(inflect.Camelize(inflect.Pluralize(builder.SchemaJSON["title"].(string))))
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
		msg :=  fmt.Sprintf("error formatting: %s, was formatting\n%s", err, src)
		b.Errors = append(b.Errors, msg)
	}
	return formatted, err
}

func (b *Builder) addHeadInfo() string {
	return "//auto-created by generator all changes will be lost.\n" +
		"//\tgo run generator.go -name=Contact -input=contact.json -o=contact.go -pkg=contacts\n\n"
}
