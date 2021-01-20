package expr

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
// pkg/expr/expr_types.yaml

import (
	"io"
	"time"
)

// Any is an expression type, wrapper for interface{} type
type Any struct{ value interface{} }

// NewAny creates new instance of Any expression type
func NewAny(new interface{}) (TypedValue, error) {
	t := &Any{}
	return t, t.Set(new)
}

// Returns underlying value on Any
func (t Any) Get() interface{} { return t.value }

// Returns type name
func (Any) Type() string { return "Any" }

// Casts value to interface{}
func (Any) Cast(value interface{}) (TypedValue, error) { return NewAny(value) }

// Boolean is an expression type, wrapper for bool type
type Boolean struct{ value bool }

// NewBoolean creates new instance of Boolean expression type
func NewBoolean(new interface{}) (TypedValue, error) {
	t := &Boolean{}
	return t, t.Set(new)
}

// Returns underlying value on Boolean
func (t Boolean) Get() interface{} { return t.value }

// Returns type name
func (Boolean) Type() string { return "Boolean" }

// Casts value to bool
func (Boolean) Cast(value interface{}) (TypedValue, error) { return NewBoolean(value) }

// DateTime is an expression type, wrapper for *time.Time type
type DateTime struct{ value *time.Time }

// NewDateTime creates new instance of DateTime expression type
func NewDateTime(new interface{}) (TypedValue, error) {
	t := &DateTime{}
	return t, t.Set(new)
}

// Returns underlying value on DateTime
func (t DateTime) Get() interface{} { return t.value }

// Returns type name
func (DateTime) Type() string { return "DateTime" }

// Casts value to *time.Time
func (DateTime) Cast(value interface{}) (TypedValue, error) { return NewDateTime(value) }

// Duration is an expression type, wrapper for time.Duration type
type Duration struct{ value time.Duration }

// NewDuration creates new instance of Duration expression type
func NewDuration(new interface{}) (TypedValue, error) {
	t := &Duration{}
	return t, t.Set(new)
}

// Returns underlying value on Duration
func (t Duration) Get() interface{} { return t.value }

// Returns type name
func (Duration) Type() string { return "Duration" }

// Casts value to time.Duration
func (Duration) Cast(value interface{}) (TypedValue, error) { return NewDuration(value) }

// Float is an expression type, wrapper for float64 type
type Float struct{ value float64 }

// NewFloat creates new instance of Float expression type
func NewFloat(new interface{}) (TypedValue, error) {
	t := &Float{}
	return t, t.Set(new)
}

// Returns underlying value on Float
func (t Float) Get() interface{} { return t.value }

// Returns type name
func (Float) Type() string { return "Float" }

// Casts value to float64
func (Float) Cast(value interface{}) (TypedValue, error) { return NewFloat(value) }

// Handle is an expression type, wrapper for string type
type Handle struct{ value string }

// NewHandle creates new instance of Handle expression type
func NewHandle(new interface{}) (TypedValue, error) {
	t := &Handle{}
	return t, t.Set(new)
}

// Returns underlying value on Handle
func (t Handle) Get() interface{} { return t.value }

// Returns type name
func (Handle) Type() string { return "Handle" }

// Casts value to string
func (Handle) Cast(value interface{}) (TypedValue, error) { return NewHandle(value) }

// ID is an expression type, wrapper for uint64 type
type ID struct{ value uint64 }

// NewID creates new instance of ID expression type
func NewID(new interface{}) (TypedValue, error) {
	t := &ID{}
	return t, t.Set(new)
}

// Returns underlying value on ID
func (t ID) Get() interface{} { return t.value }

// Returns type name
func (ID) Type() string { return "ID" }

// Casts value to uint64
func (ID) Cast(value interface{}) (TypedValue, error) { return NewID(value) }

// Integer is an expression type, wrapper for int64 type
type Integer struct{ value int64 }

// NewInteger creates new instance of Integer expression type
func NewInteger(new interface{}) (TypedValue, error) {
	t := &Integer{}
	return t, t.Set(new)
}

// Returns underlying value on Integer
func (t Integer) Get() interface{} { return t.value }

// Returns type name
func (Integer) Type() string { return "Integer" }

// Casts value to int64
func (Integer) Cast(value interface{}) (TypedValue, error) { return NewInteger(value) }

// KV is an expression type, wrapper for map[string]string type
type KV struct{ value map[string]string }

// NewKV creates new instance of KV expression type
func NewKV(new interface{}) (TypedValue, error) {
	t := &KV{}
	return t, t.Set(new)
}

// Returns underlying value on KV
func (t KV) Get() interface{} { return t.value }

// Returns type name
func (KV) Type() string { return "KV" }

// Casts value to map[string]string
func (KV) Cast(value interface{}) (TypedValue, error) { return NewKV(value) }

// KVV is an expression type, wrapper for map[string][]string type
type KVV struct{ value map[string][]string }

// NewKVV creates new instance of KVV expression type
func NewKVV(new interface{}) (TypedValue, error) {
	t := &KVV{}
	return t, t.Set(new)
}

// Returns underlying value on KVV
func (t KVV) Get() interface{} { return t.value }

// Returns type name
func (KVV) Type() string { return "KVV" }

// Casts value to map[string][]string
func (KVV) Cast(value interface{}) (TypedValue, error) { return NewKVV(value) }

// Reader is an expression type, wrapper for io.Reader type
type Reader struct{ value io.Reader }

// NewReader creates new instance of Reader expression type
func NewReader(new interface{}) (TypedValue, error) {
	t := &Reader{}
	return t, t.Set(new)
}

// Returns underlying value on Reader
func (t Reader) Get() interface{} { return t.value }

// Returns type name
func (Reader) Type() string { return "Reader" }

// Casts value to io.Reader
func (Reader) Cast(value interface{}) (TypedValue, error) { return NewReader(value) }

// String is an expression type, wrapper for string type
type String struct{ value string }

// NewString creates new instance of String expression type
func NewString(new interface{}) (TypedValue, error) {
	t := &String{}
	return t, t.Set(new)
}

// Returns underlying value on String
func (t String) Get() interface{} { return t.value }

// Returns type name
func (String) Type() string { return "String" }

// Casts value to string
func (String) Cast(value interface{}) (TypedValue, error) { return NewString(value) }

// UnsignedInteger is an expression type, wrapper for uint64 type
type UnsignedInteger struct{ value uint64 }

// NewUnsignedInteger creates new instance of UnsignedInteger expression type
func NewUnsignedInteger(new interface{}) (TypedValue, error) {
	t := &UnsignedInteger{}
	return t, t.Set(new)
}

// Returns underlying value on UnsignedInteger
func (t UnsignedInteger) Get() interface{} { return t.value }

// Returns type name
func (UnsignedInteger) Type() string { return "UnsignedInteger" }

// Casts value to uint64
func (UnsignedInteger) Cast(value interface{}) (TypedValue, error) { return NewUnsignedInteger(value) }
