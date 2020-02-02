package main

import (
	"bytes"
	"encoding/json"
	"fmt"

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
	CheckErr(err)

	errResponseSchema, _, err := openapi3gen.NewSchemaRefForValue(&api.ErrResponse{})
	CheckErr(err)

	components.Schemas = make(map[string]*openapi3.SchemaRef)
	components.Schemas["v1.User"] = userSchema
	components.Schemas["ErrResponse"] = errResponseSchema

	b := &bytes.Buffer{}
	err = json.NewEncoder(b).Encode(components.Schemas)
	CheckErr(err)

	y, err := yaml.JSONToYAML(b.Bytes())
	CheckErr(err)

	fmt.Println(string(y))
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
