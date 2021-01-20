package expr

import (
	"fmt"
	"github.com/cortezaproject/corteza-server/pkg/handle"
	"github.com/spf13/cast"
	"io"
	"net/http"
	"net/url"
	"time"
)

// Unresolved is a special type that holds value + type it needs to be resolved to
//
// This solves problem with typed value serialization
type Unresolved struct {
	typ   string
	value interface{}
}

// NewUnresolved creates new instance of Unresolved expression type
func NewUnresolved(typ string, val interface{}) (TypedValue, error) {
	return &Unresolved{
		typ:   typ,
		value: UnwindTyped(val),
	}, nil
}

// Returns underlying value on Unresolved
func (t Unresolved) Get() interface{} { return t.value }

// Returns type name
func (t Unresolved) Type() string { return t.typ }

// Casts value to interface{}
func (Unresolved) Cast(value interface{}) (TypedValue, error) {
	return nil, fmt.Errorf("can not cast unresolved")
}

func (t *Unresolved) Set(new interface{}, pp ...string) (err error) {
	return fmt.Errorf("can not set on unresolved type")
}

func (t *Any) Set(new interface{}, pp ...string) (err error) {
	if err = ReqNoPath(t.Type(), pp); err != nil {
		return
	}

	t.value = UnwindTyped(new)
	return
}

func (t *Boolean) Set(new interface{}, pp ...string) (err error) {
	var aux bool
	if err = ReqNoPath(t.Type(), pp); err != nil {
		return
	}

	if tv, is := new.(TypedValue); is {
		new = tv.Get()
	}

	if aux, err = cast.ToBoolE(new); err != nil {
		return
	}

	t.value = aux
	return
}

func (t *ID) Set(new interface{}, pp ...string) (err error) {
	return SetIDWithPath(&t.value, new, pp...)
}

func (t *Integer) Set(new interface{}, pp ...string) (err error) {
	new = UnwindTyped(new)

	var aux int64
	if err = ReqNoPath(t.Type(), pp); err != nil {
		return
	}

	if tv, is := new.(TypedValue); is {
		new = tv.Get()
	}

	if aux, err = cast.ToInt64E(new); err != nil {
		return
	}

	t.value = aux
	return
}

func (t *UnsignedInteger) Set(new interface{}, pp ...string) (err error) {
	new = UnwindTyped(new)

	var aux uint64
	if err = ReqNoPath(t.Type(), pp); err != nil {
		return
	}

	if tv, is := new.(TypedValue); is {
		new = tv.Get()
	}

	if aux, err = cast.ToUint64E(new); err != nil {
		return
	}

	t.value = aux
	return
}

func (t *Float) Set(new interface{}, pp ...string) (err error) {
	new = UnwindTyped(new)

	var aux float64
	if err = ReqNoPath(t.Type(), pp); err != nil {
		return
	}

	if aux, err = cast.ToFloat64E(new); err != nil {
		return
	}

	t.value = aux
	return
}

func (t *String) Set(new interface{}, pp ...string) (err error) {
	new = UnwindTyped(new)

	var aux string
	if err = ReqNoPath(t.Type(), pp); err != nil {
		panic(err.Error())
		return
	}

	if aux, err = cast.ToStringE(new); err != nil {
		return
	}

	t.value = aux
	return
}

func (t *Handle) Set(new interface{}, pp ...string) (err error) {
	new = UnwindTyped(new)

	var aux string
	if err = ReqNoPath(t.Type(), pp); err != nil {
		panic(err.Error())
		return
	}

	if aux, err = cast.ToStringE(new); err != nil {
		return
	}

	if !handle.IsValid(aux) {
		return fmt.Errorf("invalid handle: %q", aux)
	}

	t.value = aux
	return
}

func (t *DateTime) Set(new interface{}, pp ...string) (err error) {
	new = UnwindTyped(new)

	var aux time.Time
	if err = ReqNoPath(t.Type(), pp); err != nil {
		return
	}

	if aux, err = cast.ToTimeE(new); err != nil {
		return
	}

	t.value = &aux
	return
}

func (t *Duration) Set(new interface{}, pp ...string) (err error) {
	var aux time.Duration
	if err = ReqNoPath(t.Type(), pp); err != nil {
		return
	}

	if tv, is := new.(TypedValue); is {
		new = tv.Get()
	}

	if aux, err = cast.ToDurationE(new); err != nil {
		return
	}

	t.value = aux
	return
}

func (t *KV) Set(new interface{}, pp ...string) (err error) {
	new = UnwindTyped(new)

	if t.value == nil {
		t.value = make(map[string]string)
	}

	return SetKVWithPath(&t.value, new, pp...)
}

func (t *KVV) Set(new interface{}, pp ...string) (err error) {
	new = UnwindTyped(new)

	if t.value == nil {
		t.value = make(map[string][]string)
	}

	switch len(pp) {
	case 0:
		var aux map[string][]string
		switch casted := new.(type) {
		case KVV:
			aux = casted.value
		case *KVV:
			aux = casted.value
		case http.Header:
			aux = casted
		case url.Values:
			aux = casted
		default:
			aux, err = cast.ToStringMapStringSliceE(new)
			if err != nil {
				return err
			}
		}

		t.value = aux
	case 1:
		tmp, err := cast.ToStringSliceE(new)
		if err != nil {
			return err
		}

		t.value[pp[0]] = tmp
	default:
		return fmt.Errorf("can not set values to KVV with path deeper than 1 level")
	}

	return nil
}

func (t *Reader) Set(new interface{}, pp ...string) error {
	new = UnwindTyped(new)

	if err := ReqNoPath(t.Type(), pp); err != nil {
		return err
	}

	var ok bool
	t.value, ok = new.(io.Reader)
	if !ok {
		return fmt.Errorf("unable to cast %#v of type %T to io.Reader", new, new)
	}

	return nil
}
