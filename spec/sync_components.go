package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/aca/go-restapi-boilerplate/api"
	"github.com/aca/go-restapi-boilerplate/ent"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3gen"
	"github.com/ghodss/yaml"
)

// generate openapiv3 components from go struct
func main() {
	components := openapi3.NewComponents()

	// User
	userSchema, _, err := openapi3gen.NewSchemaRefForValue(&ent.User{})
	checkErr(err)

	errResponseSchema, _, err := openapi3gen.NewSchemaRefForValue(&api.ErrResponse{})
	checkErr(err)

	components.Schemas = make(map[string]*openapi3.SchemaRef)
	components.Schemas["v1.User"] = userSchema
	components.Schemas["ErrResponse"] = errResponseSchema

	type Swagger struct {
		Components openapi3.Components `json:"components,omitempty" yaml:"components,omitempty"`
	}

	swagger := Swagger{}
	swagger.Components = components

	b := &bytes.Buffer{}
	err = json.NewEncoder(b).Encode(swagger)
	checkErr(err)

	schema, err := yaml.JSONToYAML(b.Bytes())
	checkErr(err)

	paths, err := ioutil.ReadFile("./path.yaml")

	b = &bytes.Buffer{}
	b.Write(schema)
	b.Write(paths)

	_, err = openapi3.NewSwaggerLoader().LoadSwaggerFromData(b.Bytes())
	checkErr(err)

	err = ioutil.WriteFile("swagger.yaml", b.Bytes(), 0666)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
