// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package mixinid

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the mixinid type in the database.
	Label = "mixin_id"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSomeField holds the string denoting the some_field field in the database.
	FieldSomeField = "some_field"
	// FieldMixinField holds the string denoting the mixin_field field in the database.
	FieldMixinField = "mixin_field"
	// Table holds the table name of the mixinid in the database.
	Table = "mixin_ids"
)

// Columns holds all SQL columns for mixinid fields.
var Columns = []string{
	FieldID,
	FieldSomeField,
	FieldMixinField,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
