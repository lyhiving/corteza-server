package expr

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// extract typed-value

func TestVars_Decode(t *testing.T) {

	t.Run("mix", func(t *testing.T) {
		var (
			req = require.New(t)

			dst = &struct {
				Int        int64
				Uint64     uint64
				String     string `var:"STRING"`
				RawString  string `var:"rawString"`
				Bool       bool
				Unexisting byte
			}{}

			vars = Vars{
				"int":     Must(NewInteger(42)),
				"STRING":  Must(NewString("foo")),
				"bool":    Must(NewBoolean(true)),
				"missing": Must(NewBoolean(true)),
			}
		)

		req.NoError(vars.Decode(dst))
		req.Equal(int64(42), dst.Int)
		req.Equal("foo", dst.String)
		req.Equal(true, dst.Bool)
		req.Empty(dst.Unexisting)
	})

	t.Run("interfaces", func(t *testing.T) {
		var (
			req = require.New(t)

			dst = &struct {
				IString  interface{} `var:"iString"`
				IInteger interface{} `var:"iInteger"`
				IBool    interface{} `var:"iBool"`
			}{}

			vars = Vars{
				"iString":  Must(NewString("foo")),
				"iInteger": Must(NewInteger(42)),
				"iBool":    Must(NewBoolean(true)),
			}
		)

		req.NoError(vars.Decode(dst))
	})

	t.Run("int-uint", func(t *testing.T) {
		var (
			req = require.New(t)

			dst = &struct {
				Int    int64
				Uint64 uint64
			}{}

			vars = Vars{
				"uint64": Must(NewAny("42")),
				"int":    Must(NewAny("42")),
			}
		)

		dst.Uint64 = 0
		dst.Int = 0

		req.NoError(vars.Decode(dst))
		req.Equal(uint64(42), dst.Uint64)
		req.Equal(int64(42), dst.Int)
	})
}

func TestVars_Set(t *testing.T) {
	var (
		req = require.New(t)

		vars = Vars{
			"int": Must(NewInteger(42)),
			"sub": &Vars{
				"foo": Must(NewString("foo")),
			},
		}
	)

	req.NoError(vars.Set(Must(NewInteger(123)), "int"))
	req.Equal(int64(123), vars["int"].(TypedValue).Get().(int64))

	req.NoError(vars.Set(Must(NewString("bar")), "sub", "foo"))
	req.Equal("bar", (*(vars["sub"]).(*Vars))["foo"].Get().(string))

	req.NoError(vars.Set(&KV{}, "kv"))
	req.NoError(vars.Set(Must(NewString("bar")), "kv", "foo"))
}
