// gojson generates go struct defintions from JSON documents
//
// Reads from stdin and prints to stdout
//
// Example:
//  go run generator.go -name=Contact -input=contact.json -o=contact.go -pkg=contacts
//
// 	curl -s https://api.github.com/repos/chimeracoder/gojson | gojson -name=Repository
//

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	schema_tools "github.com/schorsch/go-json-schema-tools"
	"path"
)

var (
	name       = flag.String("name", "Foo", "the name of the struct")
	pkg        = flag.String("pkg", "main", "the name of the package for the generated code")
	inputName  = flag.String("input", "", "the name of the input file containing JSON (if input not provided via STDIN)")
	outputName = flag.String("o", "", "the name of the input file containing JSON (if input not provided via STDIN)")
)
// Return true if os.Stdin appears to be interactive
func isInteractive() bool {
	fileInfo, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return fileInfo.Mode()&(os.ModeCharDevice|os.ModeCharDevice) != 0
}

func main() {
	flag.Parse()

	if isInteractive() && *inputName == "" {
		flag.Usage()
		fmt.Fprintln(os.Stderr, "Expects input on stdin")
		os.Exit(1)
	}
	// if input_name has / use this as path,
	// else join with current working dir
	dir,_ := os.Getwd()
	file_path := path.Join(dir, *inputName)
	builder := schema_tools.NewBuilder(file_path, *name, *pkg)
	if len(builder.Errors) > 0{
		//TODO revise error outpur
		fmt.Println(builder.Errors)
		os.Exit(1)
	}
	if output, err := builder.Generate(); err != nil {
		fmt.Fprintln(os.Stderr, "error parsing ", err)
		os.Exit(1)
	} else {
		if *outputName != "" {
			err := ioutil.WriteFile(*outputName, output, 0644)
			if err != nil {
				log.Fatalf("writing output: %s", err)
			}
		} else {
			fmt.Print(string(output))
		}
	}

}
