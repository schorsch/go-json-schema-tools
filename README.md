# JSON Schema tools golang
[![Build Status](https://travis-ci.org/schorsch/go-json-schema-tools.svg?branch=master)](https://travis-ci.org/schorsch/go-json-schema-tools)

Provides a generator to create golang structs from json schema object definitions.
Still pretty raw with room for improvements.

## Usage

Checkout the repo and create a test output with:
    
    #                             # struct name  # input file                   # output file             # package name    
    go run generator/generator.go -name=Contact -input=test_schema/contact.json -o=test_schema/contact.go -pkg=contacts

To infer the package and struct name from the JSON schema title-field just omit the settings

    go run generator/generator.go -input=test_schema/contact.json -o=test_schema/contact.go
    

## Test

    go test
    # with coverage
    go test -coverprofile=coverage.out 
    go tool cover -html=coverage.out  
    
    
MIT License | Copyright 2015 Georg Leciejewski