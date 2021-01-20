package types

import (
	"fmt"
	"github.com/cortezaproject/corteza-server/pkg/expr"
	"strings"
)

type (
	ParamSet []*Param
	Param    struct {
		Name     string     `json:"name,omitempty"`
		Types    []string   `json:"types,omitempty"`
		Required bool       `json:"required,omitempty"`
		SetOf    bool       `json:"setOf,omitempty"`
		Meta     *ParamMeta `json:"meta,omitempty"`
	}

	ParamMeta struct {
		Label       string                 `json:"label,omitempty"`
		Description string                 `json:"description,omitempty"`
		Visual      map[string]interface{} `json:"visual,omitempty"`
	}

	paramOpt func(p *Param)
)

//const
func NewParam(name string, opts ...paramOpt) *Param {
	p := &Param{Name: name}
	for _, opt := range opts {
		opt(p)
	}

	return p
}

func Required(p *Param) { p.Required = !p.Required }
func SetOf(p *Param)    { p.SetOf = !p.SetOf }

func Types(tt ...expr.Type) paramOpt {
	return func(p *Param) {
		for _, t := range tt {
			p.Types = append(p.Types, t.Type())
		}
	}
}

func (p Param) HasType(t string) bool {
	for i := range p.Types {
		if p.Types[i] == t {
			return true
		}
	}
	return false
}

func (set ParamSet) GetByName(name string) *Param {
	for _, p := range set {
		if p.Name == name {
			return p
		}
	}
	return nil
}

func (set ParamSet) verify(ee ExprSet) error {
	for _, e := range ee {
		if set.GetByName(e.Target) == nil {
			return fmt.Errorf("unknown parameter %s is used", e.Target)

		}
	}

	return nil
}

// CheckArguments validates (at compile-time) input data (arguments)
func (set ParamSet) VerifyArguments(ee ExprSet) error {
	if err := set.verify(ee); err != nil {
		return err
	}

	for _, p := range set {
		e := ee.GetByTarget(p.Name)

		if e == nil {
			if p.Required {
				return fmt.Errorf("parameter %s is required", p.Name)
			}

			continue
		}

		if !p.HasType(e.Type) && !p.HasType(expr.Any{}.Type()) {
			return fmt.Errorf(
				"incompatible argument type %s for parameter %s, expecting %s",
				e.Type, p.Name,
				strings.Join(p.Types, ", "),
			)
		}

		// @todo check if target holds set-of values (array of values)
		//       this could be implemented by generic wrapping type that would
		//       enable
	}

	return nil
}

// CheckArguments validates (at compile-time) input data (arguments)
func (set ParamSet) VerifyResults(ee ExprSet) error {
	if err := set.verify(ee); err != nil {
		return err
	}

	for _, p := range set {
		e := ee.GetByTarget(p.Name)
		if e == nil {
			continue
		}

		if e.Type != "" && !p.HasType(e.Type) {
			return fmt.Errorf("incompatible type %s for result %s, expecting %s",
				e.Type, p.Name,
				strings.Join(p.Types, ", "),
			)
		}
	}

	return nil
}
