package automation

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
// compose/automation/expr_types.yaml

import (
	"github.com/cortezaproject/corteza-server/compose/types"
	"github.com/cortezaproject/corteza-server/pkg/expr"
)

// Module is an expression type, wrapper for *types.Module type
type Module struct{ value *types.Module }

// NewModule creates new instance of Module expression type
func NewModule(new interface{}) (expr.TypedValue, error) {
	t := &Module{}
	return t, t.Set(new)
}

// Returns underlying value on Module
func (t Module) Get() interface{} { return t.value }

// Returns type name
func (Module) Type() string { return "ComposeModule" }

// Casts value to *types.Module
func (Module) Cast(value interface{}) (expr.TypedValue, error) { return NewModule(value) }

// Namespace is an expression type, wrapper for *types.Namespace type
type Namespace struct{ value *types.Namespace }

// NewNamespace creates new instance of Namespace expression type
func NewNamespace(new interface{}) (expr.TypedValue, error) {
	t := &Namespace{}
	return t, t.Set(new)
}

// Returns underlying value on Namespace
func (t Namespace) Get() interface{} { return t.value }

// Returns type name
func (Namespace) Type() string { return "ComposeNamespace" }

// Casts value to *types.Namespace
func (Namespace) Cast(value interface{}) (expr.TypedValue, error) { return NewNamespace(value) }

// Record is an expression type, wrapper for *types.Record type
type Record struct{ value *types.Record }

// NewRecord creates new instance of Record expression type
func NewRecord(new interface{}) (expr.TypedValue, error) {
	t := &Record{}
	return t, t.Set(new)
}

// Returns underlying value on Record
func (t Record) Get() interface{} { return t.value }

// Returns type name
func (Record) Type() string { return "ComposeRecord" }

// Casts value to *types.Record
func (Record) Cast(value interface{}) (expr.TypedValue, error) { return NewRecord(value) }

// RecordValueErrorSet is an expression type, wrapper for *types.RecordValueErrorSet type
type RecordValueErrorSet struct{ value *types.RecordValueErrorSet }

// NewRecordValueErrorSet creates new instance of RecordValueErrorSet expression type
func NewRecordValueErrorSet(new interface{}) (expr.TypedValue, error) {
	t := &RecordValueErrorSet{}
	return t, t.Set(new)
}

// Returns underlying value on RecordValueErrorSet
func (t RecordValueErrorSet) Get() interface{} { return t.value }

// Returns type name
func (RecordValueErrorSet) Type() string { return "ComposeRecordValueErrorSet" }

// Casts value to *types.RecordValueErrorSet
func (RecordValueErrorSet) Cast(value interface{}) (expr.TypedValue, error) {
	return NewRecordValueErrorSet(value)
}

// RecordValues is an expression type, wrapper for types.RecordValueSet type
type RecordValues struct{ value types.RecordValueSet }

// NewRecordValues creates new instance of RecordValues expression type
func NewRecordValues(new interface{}) (expr.TypedValue, error) {
	t := &RecordValues{}
	return t, t.Set(new)
}

// Returns underlying value on RecordValues
func (t RecordValues) Get() interface{} { return t.value }

// Returns type name
func (RecordValues) Type() string { return "ComposeRecordValues" }

// Casts value to types.RecordValueSet
func (RecordValues) Cast(value interface{}) (expr.TypedValue, error) { return NewRecordValues(value) }
