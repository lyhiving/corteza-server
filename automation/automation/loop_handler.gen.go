package automation

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
// automation/automation/loop_handler.yaml

import (
	"context"
	atypes "github.com/cortezaproject/corteza-server/automation/types"
	"github.com/cortezaproject/corteza-server/pkg/expr"
	"github.com/cortezaproject/corteza-server/pkg/logger"
	"github.com/cortezaproject/corteza-server/pkg/wfexec"
	"go.uber.org/zap"
	"io"
)

var _ wfexec.ExecResponse

type (
	loopHandlerRegistry interface {
		AddFunctions(ff ...*atypes.Function)
		Type(ref string) expr.Type
	}
)

func (h loopHandler) register() {
	h.reg.AddFunctions(
		h.Sequence(),
		h.Do(),
		h.Each(),
		h.Lines(),
	)
}

type (
	loopSequenceArgs struct {
		log *zap.Logger

		hasFirst bool

		First int64

		hasLast bool

		Last int64

		hasStep bool

		Step int64
	}
)

//
//
// expects implementation of sequence function:
// func (h loop) sequence(ctx context.Context, args *loopSequenceArgs) (results *loopSequenceResults, err error) {
//    return
// }
func (h loopHandler) Sequence() *atypes.Function {
	return &atypes.Function{
		Ref: "loopSequence",
		Meta: &atypes.FunctionMeta{
			Short: "Iterates over sequence of numbers",
		},

		Parameters: []*atypes.Param{
			{
				Name:  "first",
				Types: []string{"Integer"},
			},
			{
				Name:  "last",
				Types: []string{"Integer"},
			},
			{
				Name:  "step",
				Types: []string{"Integer"},
			},
		},

		IsIterator: true,

		Iterator: func(ctx context.Context, in expr.Vars) (out wfexec.IteratorHandler, err error) {
			var (
				args = &loopSequenceArgs{
					log: logger.ContextValue(ctx, zap.NewNop()).
						With(zap.String("function", "loopSequence")),
					hasFirst: in.Has("first"),
					hasLast:  in.Has("last"),
					hasStep:  in.Has("step"),
				}
			)

			if err = in.Decode(args); err != nil {
				return
			}

			return h.sequence(ctx, args)
		},
	}
}

type (
	loopDoArgs struct {
		log *zap.Logger

		hasWhile bool

		While string
	}
)

//
//
// expects implementation of do function:
// func (h loop) do(ctx context.Context, args *loopDoArgs) (results *loopDoResults, err error) {
//    return
// }
func (h loopHandler) Do() *atypes.Function {
	return &atypes.Function{
		Ref: "loopDo",
		Meta: &atypes.FunctionMeta{
			Short: "Iterates while condition is true",
		},

		Parameters: []*atypes.Param{
			{
				Name:  "while",
				Types: []string{"String"}, Required: true,
				Meta: &atypes.ParamMeta{
					Label:       "Expression tested before each iteration",
					Description: "Expression to be evaluated each iteration; loop will continue until expression is true",
				},
			},
		},

		IsIterator: true,

		Iterator: func(ctx context.Context, in expr.Vars) (out wfexec.IteratorHandler, err error) {
			var (
				args = &loopDoArgs{
					log: logger.ContextValue(ctx, zap.NewNop()).
						With(zap.String("function", "loopDo")),
					hasWhile: in.Has("while"),
				}
			)

			if err = in.Decode(args); err != nil {
				return
			}

			return h.do(ctx, args)
		},
	}
}

type (
	loopEachArgs struct {
		log *zap.Logger

		hasItems bool

		Items []interface{}
	}
)

//
//
// expects implementation of each function:
// func (h loop) each(ctx context.Context, args *loopEachArgs) (results *loopEachResults, err error) {
//    return
// }
func (h loopHandler) Each() *atypes.Function {
	return &atypes.Function{
		Ref: "loopEach",
		Meta: &atypes.FunctionMeta{
			Short: "Iterates over set of items",
		},

		Parameters: []*atypes.Param{
			{
				Name:  "items",
				Types: []string{"Any"}, Required: true,
			},
		},

		IsIterator: true,

		Iterator: func(ctx context.Context, in expr.Vars) (out wfexec.IteratorHandler, err error) {
			var (
				args = &loopEachArgs{
					log: logger.ContextValue(ctx, zap.NewNop()).
						With(zap.String("function", "loopEach")),
					hasItems: in.Has("items"),
				}
			)

			if err = in.Decode(args); err != nil {
				return
			}

			return h.each(ctx, args)
		},
	}
}

type (
	loopLinesArgs struct {
		log *zap.Logger

		hasStream bool

		Stream io.Reader
	}
)

//
//
// expects implementation of lines function:
// func (h loop) lines(ctx context.Context, args *loopLinesArgs) (results *loopLinesResults, err error) {
//    return
// }
func (h loopHandler) Lines() *atypes.Function {
	return &atypes.Function{
		Ref: "loopLines",
		Meta: &atypes.FunctionMeta{
			Short: "Iterates over lines from stream",
		},

		Parameters: []*atypes.Param{
			{
				Name:  "stream",
				Types: []string{"Reader"}, Required: true,
			},
		},

		IsIterator: true,

		Iterator: func(ctx context.Context, in expr.Vars) (out wfexec.IteratorHandler, err error) {
			var (
				args = &loopLinesArgs{
					log: logger.ContextValue(ctx, zap.NewNop()).
						With(zap.String("function", "loopLines")),
					hasStream: in.Has("stream"),
				}
			)

			if err = in.Decode(args); err != nil {
				return
			}

			return h.lines(ctx, args)
		},
	}
}
