package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
	"gopkg.in/yaml.v3"
)

func main() {
	sources, errRead := readSchemas(schemasFlag)
	if errRead != nil {
		panic(errRead)
	}

	schema, errLoad := loadSchema(sources)
	if errLoad != nil {
		panic(errLoad)
	}

	marshal := saveToYAML
	if asJSON {
		marshal = saveToJSON
	}
	if err := marshal(resultFile, schema); err != nil {
		panic(err)
	}
}

func readSchemas(schemas []string) ([]*ast.Source, error) {
	sources := make([]*ast.Source, 0, len(schemas))
	for _, filename := range schemas {
		f, err := os.Open(filename)
		if err != nil {
			return nil, fmt.Errorf("cannot open file %q: %w", filename, err)
		}
		defer f.Close()

		schemaRaw, err := io.ReadAll(f)
		if err != nil {
			return nil, fmt.Errorf("cannot read file %q: %w", filename, err)
		}

		sources = append(sources, &ast.Source{
			Name:  filename,
			Input: string(schemaRaw),
		})
	}
	return sources, nil
}

func loadSchema(sources []*ast.Source) (*ast.Schema, error) {
	schema, err := gqlparser.LoadSchema(sources...)
	if err != nil {
		return nil, fmt.Errorf("cannot load schema: %w", err)
	}

	if schema.Query == nil {
		schema.Query = &ast.Definition{
			Kind: ast.Object,
			Name: "Query",
		}
		schema.Types["Query"] = schema.Query
	}
	return schema, nil
}

var schemasFlag stringSliceFlag
var resultFile string
var asJSON bool

func init() {
	flag.Var(&schemasFlag, "schema", "file with a GraphQL schema")
	flag.StringVar(&resultFile, "result", "schema.yaml", "filename with a schema in YAML")
	flag.BoolVar(&asJSON, "json", false, "should save schema in JSON")
	flag.Parse()

	if len(schemasFlag) == 0 {
		panic("no schema files provided")
	}
	if asJSON && resultFile == "schema.yaml" {
		resultFile = "schema.json"
	}
}

func saveToYAML(filaname string, v interface{}) error {
	result, err := os.Create(filaname)
	if err != nil {
		return fmt.Errorf("cannot open result file: %w", err)
	}
	defer result.Close()

	if err := yaml.NewEncoder(result).Encode(v); err != nil {
		return fmt.Errorf("cannot marshal to YAML: %w", err)
	}
	return nil
}

func saveToJSON(filaname string, v interface{}) error {
	result, err := os.Create(filaname)
	if err != nil {
		return fmt.Errorf("cannot open result file: %w", err)
	}
	defer result.Close()

	if err := json.NewEncoder(result).Encode(v); err != nil {
		return fmt.Errorf("cannot marshal to JSON: %w", err)
	}
	return nil
}

type stringSliceFlag []string

func (ss *stringSliceFlag) String() string {
	return strings.Join(*ss, ",")
}

func (ss *stringSliceFlag) Set(value string) error {
	*ss = append(*ss, value)
	return nil
}
