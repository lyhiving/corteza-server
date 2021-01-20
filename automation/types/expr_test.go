package types

import (
	"context"
	"github.com/cortezaproject/corteza-server/pkg/expr"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestExprSet_Eval(t *testing.T) {
	var (
		ctx = context.Background()

		// extract typed-value
		etv = func(v expr.TypedValue, err error) expr.TypedValue {
			if err != nil {
				panic(err)
			}
			return v
		}

		cc = []struct {
			name   string
			set    ExprSet
			input  expr.Vars
			output expr.Vars
			err    string
		}{
			{
				name:   "empty",
				set:    ExprSet{},
				output: expr.Vars{},
			},
			{
				name:   "constant assignment",
				set:    ExprSet{&Expr{Target: "foo", Expr: `"bar"`}},
				output: expr.Vars{"foo": etv(expr.NewAny("bar"))},
			},
			{
				name:   "vars with path",
				set:    ExprSet{&Expr{Target: "l1.l2", Expr: `"bar"`}},
				input:  expr.Vars{"l1": &expr.Vars{}},
				output: expr.Vars{"l1": &expr.Vars{"l2": etv(expr.NewAny("bar"))}},
			},
			{
				name: "copy vars with same types",
				set: ExprSet{
					&Expr{Target: "aa", Value: "vv", typ: &expr.String{}},
					&Expr{Target: "bb", Source: "aa", typ: &expr.String{}},
				},
				output: expr.Vars{
					"aa": etv(expr.NewString("vv")),
					"bb": etv(expr.NewString("vv")),
				},
			},
			{
				name: "copy var with type",
				set: ExprSet{
					&Expr{Target: "aa", Value: "should be always String", typ: &expr.String{}},
					&Expr{Target: "bb", Source: "aa"},
				},
				output: expr.Vars{
					"aa": etv(expr.NewString("should be always String")),
					"bb": etv(expr.NewString("should be always String")),
				},
			},
			{
				name: "copy var to target with type",
				set: ExprSet{
					&Expr{Target: "aa", Value: "42", typ: &expr.String{}},
					&Expr{Target: "bb", Source: "aa", typ: &expr.Integer{}},
				},
				output: expr.Vars{
					"aa": etv(expr.NewString("42")),
					"bb": etv(expr.NewInteger(42)),
				},
			},
			{
				name: "assign into incompatible",
				set: ExprSet{
					&Expr{Target: "aa", Value: "foo", typ: &expr.String{}},
					&Expr{Target: "bb", Source: "aa", typ: &expr.Integer{}},
				},
				err: "unable to cast \"foo\" of type string to int64",
			},
		}
	)
	for _, c := range cc {
		t.Run(c.name, func(t *testing.T) {
			var (
				req = require.New(t)
			)

			for _, e := range c.set {
				if e.Expr != "" {
					req.NoError(expr.NewGvalParser().ParseEvaluators(e))
				}

				if e.typ == nil {
					e.typ = expr.Any{}
				}
			}

			var (
				output, err = c.set.Eval(ctx, c.input)
			)

			if c.err == "" {
				req.NoError(err)
			} else {
				req.Error(err, c.err)
			}

			req.Equal(c.output, output)
		})
	}
}
