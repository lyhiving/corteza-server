package automation

import (
	"bufio"
	"context"
	"github.com/cortezaproject/corteza-server/pkg/expr"
)

type (
	// iterates from start to stop by step
	sequenceIterator struct {
		counter, cFirst, cLast, cStep int64
	}
)

func (i *sequenceIterator) More(context.Context, expr.Vars) (bool, error) {
	return i.counter*(i.cStep/i.cStep) < i.cLast*(i.cStep/i.cStep), nil
}

func (i *sequenceIterator) Start(context.Context, expr.Vars) error { return nil }

func (i *sequenceIterator) Next(context.Context, expr.Vars) (expr.Vars, error) {
	scope := expr.Vars{}
	scope["counter"], _ = expr.NewInteger(i.counter)
	scope["first"], _ = expr.NewInteger(i.cStep)
	scope["last"], _ = expr.NewInteger(i.counter)
	scope["step"], _ = expr.NewInteger(i.counter)

	i.counter = i.counter + i.cStep
	return scope, nil
}

type (
	// iterates from start to stop by step
	conditionIterator struct {
		expr expr.Evaluable
	}
)

func (i *conditionIterator) More(ctx context.Context, scope expr.Vars) (bool, error) {
	return i.expr.Test(ctx, scope)
}

func (i *conditionIterator) Start(context.Context, expr.Vars) error { return nil }

func (i *conditionIterator) Next(context.Context, expr.Vars) (expr.Vars, error) {
	return expr.Vars{}, nil
}

type (
	// iterates from start to stop by step
	collectionIterator struct {
		ptr int
		set []interface{}
	}
)

func (i *collectionIterator) More(ctx context.Context, scope expr.Vars) (bool, error) {
	return i.ptr < len(i.set), nil
}

func (i *collectionIterator) Start(context.Context, expr.Vars) error { i.ptr = 0; return nil }

func (i *collectionIterator) Next(context.Context, expr.Vars) (expr.Vars, error) {
	if item, err := expr.NewAny(i.set[i.ptr]); err != nil {
		return nil, err
	} else {
		i.ptr++
		return expr.Vars{"item": item}, nil
	}
}

type (
	// iterates from start to stop by step
	lineIterator struct {
		s *bufio.Scanner
	}
)

func (i *lineIterator) More(context.Context, expr.Vars) (bool, error) {
	return i.s.Scan(), nil
}

func (i *lineIterator) Start(context.Context, expr.Vars) error {
	return nil
}

func (i *lineIterator) Next(context.Context, expr.Vars) (expr.Vars, error) {
	if err := i.s.Err(); err != nil {
		return nil, err
	}

	line, _ := expr.NewString(i.s.Text())
	return expr.Vars{"line": line}, nil
}
