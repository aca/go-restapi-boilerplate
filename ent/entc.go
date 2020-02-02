// +build ignore

package main

import (
	"log"

	"github.com/facebookincubator/ent/entc"
	"github.com/facebookincubator/ent/entc/gen"
	"github.com/facebookincubator/ent/schema/field"
)

func main() {
	err := entc.Generate("./schema", &gen.Config{
		Header: `
			// Code generated (@generated) by entc, DO NOT EDIT.
		`,
		IDType: &field.TypeInfo{Type: field.TypeInt},
	})
	if err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
