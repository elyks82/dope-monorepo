// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/dopedao/dope-monorepo/packages/api/ent/dope"
	"github.com/dopedao/dope-monorepo/packages/api/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	dopeFields := schema.Dope{}.Fields()
	_ = dopeFields
	// dopeDescClaimed is the schema descriptor for claimed field.
	dopeDescClaimed := dopeFields[1].Descriptor()
	// dope.DefaultClaimed holds the default value on creation for the claimed field.
	dope.DefaultClaimed = dopeDescClaimed.Default.(bool)
	// dopeDescOpened is the schema descriptor for opened field.
	dopeDescOpened := dopeFields[2].Descriptor()
	// dope.DefaultOpened holds the default value on creation for the opened field.
	dope.DefaultOpened = dopeDescOpened.Default.(bool)
}