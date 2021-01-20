package expr

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/PaesslerAG/gval"
	"github.com/cortezaproject/corteza-server/pkg/errors"
	"github.com/spf13/cast"
	"reflect"
	"strings"
)

type (
	Type interface {
		Type() string
		Cast(interface{}) (TypedValue, error)
	}

	TypedValue interface {
		Type
		Setter
		Get() interface{}
	}

	typedValueWrap struct {
		Value interface{} `json:"@value"`
		Type  string      `json:"@type"`
	}

	Setter interface {
		Set(interface{}, ...string) error
	}

	Decoder interface {
		Decode(reflect.Value) error
	}

	Dict interface {
		Dict() map[string]interface{}
	}
)

func Must(v TypedValue, err error) TypedValue {
	if err != nil {
		panic(err)
	}
	return v
}

func ReqNoPath(t string, pp []string) error {
	if len(pp) > 0 {
		return fmt.Errorf("setting values with path on %s is not supported", t)
	}

	return nil
}

type Vars map[string]TypedValue

var _ TypedValue = Vars{}

func (Vars) Type() string { return "Vars" }
func (Vars) Cast(value interface{}) (TypedValue, error) {
	switch casted := value.(type) {
	case Vars:
		return casted, nil
	default:
		return nil, fmt.Errorf("unable to cast type %T to Vars", value)
	}
}

func (vv Vars) ResolveTypes(res func(typ string) Type) (err error) {
	for k, v := range vv {
		fmt.Printf("resolving %s %T\n", k, v)
		if u, is := v.(*Unresolved); is {
			if res(u.Type()) == nil {
				return errors.NotFound("failed to resolve unknown or unregistered type %q on %q", u.Type(), k)
			}

			vv[k], err = res(u.Type()).Cast(vv[k])
			if err != nil {
				return fmt.Errorf("failed to resolve: %w", err)
			}
		}

		if r, is := vv[k].(Vars); is {
			if err = r.ResolveTypes(res); err != nil {
				return
			}
		}
	}

	return nil
}

func (vv Vars) Set(new interface{}, pp ...string) (err error) {
	var value TypedValue
	if aux, is := new.(TypedValue); is {
		value = aux
	} else if value, err = NewAny(value); err != nil {
		return err
	}

	if len(pp) > 0 {
		p := pp[0]
		if _, has := vv[p]; !has {
			if len(pp) > 1 {
				return fmt.Errorf("%q does not exist, can not set value with path %q", p, strings.Join(pp[1:], "."))
			}

			vv[p] = value
		}

		if sub, is := vv[p].(*Vars); is {
			return sub.Set(new, pp[1:]...)
		} else if sub, is := vv[p].(Setter); is {
			return sub.Set(value, pp[1:]...)
		} else {
			if len(pp) > 1 {
				return fmt.Errorf("can not set value with path %q on %T", strings.Join(pp[1:], "."), vv[p])
			} else {
				vv[p] = value
			}
		}

		return nil
	} else {
		return fmt.Errorf("unable to replace Vars")
	}
}

func (vv Vars) Get() interface{} {
	return vv
}

// Assign takes base variables and assigns all new variables
func (vv Vars) Merge(nn ...Vars) Vars {
	var (
		out = Vars{}
	)

	nn = append([]Vars{vv}, nn...)
	for i := range nn {
		for k, v := range nn[i] {
			out[k] = v
		}
	}

	return out
}

// Returns true if all keys are present
func (vv Vars) Has(key string, kk ...string) bool {
	for _, key = range append([]string{key}, kk...) {
		if _, has := vv[key]; !has {
			return false
		}
	}

	return true
}

// Returns true if all keys are present
func (vv Vars) Any(key string, kk ...string) bool {
	for _, key = range append([]string{key}, kk...) {
		if _, has := vv[key]; has {
			return true
		}
	}

	return false
}

func (vv Vars) Dict() map[string]interface{} {
	dict := make(map[string]interface{})
	for k, v := range vv {
		switch v := v.(type) {
		case gval.Selector:
			dict[k] = v

		case Dict:
			dict[k] = v.Dict()

		case TypedValue:
			dict[k] = v.Get()

		default:
			dict[k] = v
		}

	}

	return dict
}

func (vv Vars) Decode(dst interface{}) (err error) {
	dstRef := reflect.ValueOf(dst)

	if dstRef.Kind() != reflect.Ptr {
		return fmt.Errorf("expecting a pointer, not a value")
	}

	if dstRef.IsNil() {
		return fmt.Errorf("nil pointer passed")
	}

	dstRef = dstRef.Elem()

	for i := 0; i < dstRef.NumField(); i++ {

		var (
			value interface{}
			has   bool
			t     = dstRef.Type().Field(i)
		)

		keyName := t.Tag.Get("var")
		if keyName == "" {
			keyName = strings.ToLower(t.Name[:1]) + t.Name[1:]
		}

		value, has = vv[keyName]
		if !has {
			continue
		}

		if err = decode(dstRef.Field(i), value); err != nil {
			return fmt.Errorf("failed to decode value to field %s: %w", t.Name, err)
		}
	}

	return
}

func (vv *Vars) Scan(value interface{}) error {
	//lint:ignore S1034 This typecast is intentional, we need to get []byte out of a []uint8
	switch value.(type) {
	case nil:
		*vv = Vars{}
	case []uint8:
		b := value.([]byte)
		if err := json.Unmarshal(b, vv); err != nil {
			return fmt.Errorf("can not scan '%v' into %T: %w", string(b), vv, err)
		}
	}

	return nil
}

func (vv Vars) Value() (driver.Value, error) {
	return json.Marshal(vv)
}

// UnmarshalJSON
func (vv *Vars) UnmarshalJSON(in []byte) (err error) {
	var (
		aux = make(map[string]*typedValueWrap)
	)

	if *vv == nil {
		*vv = Vars{}
	}

	if err = json.Unmarshal(in, &aux); err != nil {
		return
	}

	for k, v := range aux {
		if (*vv)[k], err = NewUnresolved(v.Type, v.Value); err != nil {
			return
		}
	}

	return nil
}

// UnmarshalJSON parses sort expression when passed inside JSON
func (vv Vars) MarshalJSON() ([]byte, error) {
	aux := make(map[string]*typedValueWrap)
	for k, v := range vv {
		aux[k] = &typedValueWrap{Type: v.Type()}

		if _, is := v.(json.Marshaler); is {
			aux[k].Value = v
		} else {
			aux[k].Value = v.Get()
		}
	}

	return json.Marshal(aux)
}

func decode(dst reflect.Value, src interface{}) (err error) {
	defer func() {
		r := recover()
		if r == nil {
			return
		}

		switch r := r.(type) {
		case error:
			err = r
		default:
			err = fmt.Errorf("%v", r)
		}
	}()

	if um, is := src.(Decoder); is {
		return um.Decode(dst)
	}

	src = UnwindTyped(src)

	var (
		vBool    bool
		vInt64   int64
		vUint64  uint64
		vFloat64 float64
		vString  string
	)

	switch dst.Kind() {
	case reflect.Bool:
		if vBool, err = cast.ToBoolE(src); err == nil {
			dst.SetBool(vBool)
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if vInt64, err = cast.ToInt64E(src); err == nil {
			dst.SetInt(vInt64)
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if vUint64, err = cast.ToUint64E(src); err == nil {
			dst.SetUint(vUint64)
		}

	case reflect.Float32, reflect.Float64:
		if vFloat64, err = cast.ToFloat64E(src); err == nil {
			dst.SetFloat(vFloat64)
		}

	case reflect.String:
		if vString, err = cast.ToStringE(src); err == nil {
			dst.SetString(vString)
		}

	case reflect.Interface:
		dst.Set(reflect.ValueOf(src))

	default:
		return fmt.Errorf("failed to cast %T to %s", src, dst.Kind())
	}

	if err != nil {
		return fmt.Errorf("failed to cast %T to %s: %w", src, dst.Kind(), err)
	}

	return nil
}
