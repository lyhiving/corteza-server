package types

import (
	"context"
	"fmt"
	"github.com/cortezaproject/corteza-server/pkg/expr"
	"github.com/cortezaproject/corteza-server/pkg/wfexec"
)

type (
	FunctionHandler func(ctx context.Context, in expr.Vars) (expr.Vars, error)
	IteratorHandler func(ctx context.Context, in expr.Vars) (wfexec.IteratorHandler, error)

	// workflow functions are defined in the core code and through plugins
	Function struct {
		Ref        string        `json:"ref,omitempty"`
		IsIterator bool          `json:"iterator,omitempty"`
		Meta       *FunctionMeta `json:"meta,omitempty"`
		Parameters ParamSet      `json:"parameters,omitempty"`
		Results    ParamSet      `json:"results,omitempty"`

		Handler  FunctionHandler `json:"-"`
		Iterator IteratorHandler `json:"-"`
	}

	FunctionMeta struct {
		Short       string                 `json:"short,omitempty"`
		Description string                 `json:"description,omitempty"`
		Visual      map[string]interface{} `json:"visual,omitempty"`
	}

	functionStep struct {
		identifiableStep
		def       *Function
		arguments ExprSet
		results   ExprSet
	}

	iteratorStep struct {
		identifiableStep
		def       *Function
		arguments ExprSet
		next      wfexec.Step
		exit      wfexec.Step
	}
)

func FunctionStep(def *Function, arguments, results ExprSet) (*functionStep, error) {
	if def.IsIterator {
		return nil, fmt.Errorf("expecting non-iterator function")
	}

	return &functionStep{def: def, arguments: arguments, results: results}, nil
}

func (f *functionStep) Exec(ctx context.Context, r *wfexec.ExecRequest) (wfexec.ExecResponse, error) {
	var (
		args, results expr.Vars
		err           error
	)

	if len(f.arguments) > 0 {
		// Arguments defined, get values from scope and use them when calling
		// function/handler
		args, err = f.arguments.Eval(ctx, r.Scope.Merge(r.Input))
		if err != nil {
			return nil, err
		}
	}

	results, err = f.def.Handler(ctx, args)
	if err != nil {
		return nil, err
	}

	if len(f.results) == 0 {
		// No results defined, nothing to return
		return expr.Vars{}, nil
	}

	results, err = f.results.Eval(ctx, results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func IteratorStep(def *Function, arguments ExprSet, next, exit wfexec.Step) (*iteratorStep, error) {
	if !def.IsIterator {
		return nil, fmt.Errorf("expecting iterator function")
	}

	return &iteratorStep{def: def, arguments: arguments, next: next, exit: exit}, nil
}

func (f *iteratorStep) Exec(ctx context.Context, r *wfexec.ExecRequest) (wfexec.ExecResponse, error) {
	var (
		args expr.Vars
		err  error
		ih   wfexec.IteratorHandler
	)

	if len(f.arguments) > 0 {
		// Arguments defined, get values from scope and use them when calling
		// iterator/handler
		args, err = f.arguments.Eval(ctx, r.Scope.Merge(r.Input))
		if err != nil {
			return nil, err
		}
	}

	if ih, err = f.def.Iterator(ctx, args); err != nil {
		return nil, err
	}

	return wfexec.GenericIterator(f, f.next, f.exit, ih), nil
}
