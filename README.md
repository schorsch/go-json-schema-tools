# JSON Schema tools golang

Provides a generator to create golang structs from json schema object definitions.
Still pretty raw with room for improvements.

## Usage

Checkout the repo and create a test output with:
    
    #                             # struct name  # input file                   # output file             # package name    
    go run generator/generator.go -name=Contact -input=test_schema/contact.json -o=test_schema/contact.go -pkg=contacts


## Test

    go test
    # with coverage
    go test -coverprofile=coverage.out 
    go tool cover -html=coverage.out  
    
    
MIT License | Copyright 2015 Georg Leciejewski