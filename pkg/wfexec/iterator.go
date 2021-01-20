package wfexec

import (
	"context"
	"github.com/cortezaproject/corteza-server/pkg/expr"
)

type (
	// Iterator can be returned from Exec fn as ExecResponse
	//
	// It helps session's exec fn() to properly navigate through graph
	// by calling is/break/iterator/next function
	Iterator interface {
		// Is the given step this iterator step
		Is(Step) bool

		// Initialize iterator
		Start(context.Context, expr.Vars) error

		// Break fn is called when loop is forcefully broken
		Break() Step

		Iterator() Step

		// Next is called before each iteration and returns
		// 1st step of the iteration branch and variables that are added to the scope
		Next(context.Context, expr.Vars) (Step, expr.Vars, error)
	}

	genericIterator struct {
		iter, next, exit Step

		h IteratorHandler
	}

	IteratorHandler interface {
		Start(context.Context, expr.Vars) error
		More(context.Context, expr.Vars) (bool, error)
		Next(context.Context, expr.Vars) (expr.Vars, error)
	}
)

func GenericIterator(iter, next, exit Step, h IteratorHandler) Iterator {
	return &genericIterator{
		iter: iter,
		next: next,
		exit: exit,
		h:    h,
	}
}

func (i *genericIterator) Is(s Step) bool                               { return i.iter == s }
func (i *genericIterator) Start(ctx context.Context, s expr.Vars) error { return i.h.Start(ctx, s) }
func (i *genericIterator) Break() Step                                  { return i.exit }
func (i *genericIterator) Iterator() Step                               { return i.iter }
func (i *genericIterator) Next(ctx context.Context, scope expr.Vars) (next Step, out expr.Vars, err error) {
	var more bool
	if more, err = i.h.More(ctx, scope); err != nil || !more {
		return
	}

	next = i.next
	out, err = i.h.Next(ctx, scope)
	return
}
