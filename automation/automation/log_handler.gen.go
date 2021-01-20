package automation

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
// automation/automation/log_handler.yaml

import (
	"context"
	atypes "github.com/cortezaproject/corteza-server/automation/types"
	"github.com/cortezaproject/corteza-server/pkg/expr"
	"github.com/cortezaproject/corteza-server/pkg/wfexec"
)

var _ wfexec.ExecResponse

type (
	logHandlerRegistry interface {
		AddFunctions(ff ...*atypes.Function)
		Type(ref string) expr.Type
	}
)

func (h logHandler) register() {
	h.reg.AddFunctions(
		h.Debug(),
		h.Info(),
		h.Warn(),
		h.Error(),
	)
}

type (
	logDebugArgs struct {
		hasMessage bool
		Message    string

		hasFields bool
		Fields    map[string]string
	}
)

//
//
// expects implementation of debug function:
// func (h log) debug(ctx context.Context, args *logDebugArgs) (results *logDebugResults, err error) {
//    return
// }
func (h logHandler) Debug() *atypes.Function {
	return &atypes.Function{
		Ref:  "logDebug",
		Type: "",
		Meta: &atypes.FunctionMeta{
			Short: "Writes debug log message",
		},

		Parameters: []*atypes.Param{
			{
				Name:  "message",
				Types: []string{"String"}, Required: true,
			},
			{
				Name:  "fields",
				Types: []string{"KV"},
			},
		},

		Handler: func(ctx context.Context, in expr.Vars) (out expr.Vars, err error) {
			var (
				args = &logDebugArgs{
					hasMessage: in.Has("message"),
					hasFields:  in.Has("fields"),
				}
			)

			if err = in.Decode(args); err != nil {
				return
			}

			return out, h.debug(ctx, args)
		},
	}
}

type (
	logInfoArgs struct {
		hasMessage bool
		Message    string

		hasFields bool
		Fields    map[string]string
	}
)

//
//
// expects implementation of info function:
// func (h log) info(ctx context.Context, args *logInfoArgs) (results *logInfoResults, err error) {
//    return
// }
func (h logHandler) Info() *atypes.Function {
	return &atypes.Function{
		Ref:  "logInfo",
		Type: "",
		Meta: &atypes.FunctionMeta{
			Short: "Writes info log message",
		},

		Parameters: []*atypes.Param{
			{
				Name:  "message",
				Types: []string{"String"}, Required: true,
			},
			{
				Name:  "fields",
				Types: []string{"KV"},
			},
		},

		Handler: func(ctx context.Context, in expr.Vars) (out expr.Vars, err error) {
			var (
				args = &logInfoArgs{
					hasMessage: in.Has("message"),
					hasFields:  in.Has("fields"),
				}
			)

			if err = in.Decode(args); err != nil {
				return
			}

			return out, h.info(ctx, args)
		},
	}
}

type (
	logWarnArgs struct {
		hasMessage bool
		Message    string

		hasFields bool
		Fields    map[string]string
	}
)

//
//
// expects implementation of warn function:
// func (h log) warn(ctx context.Context, args *logWarnArgs) (results *logWarnResults, err error) {
//    return
// }
func (h logHandler) Warn() *atypes.Function {
	return &atypes.Function{
		Ref:  "logWarn",
		Type: "",
		Meta: &atypes.FunctionMeta{
			Short: "Writes warn log message",
		},

		Parameters: []*atypes.Param{
			{
				Name:  "message",
				Types: []string{"String"}, Required: true,
			},
			{
				Name:  "fields",
				Types: []string{"KV"},
			},
		},

		Handler: func(ctx context.Context, in expr.Vars) (out expr.Vars, err error) {
			var (
				args = &logWarnArgs{
					hasMessage: in.Has("message"),
					hasFields:  in.Has("fields"),
				}
			)

			if err = in.Decode(args); err != nil {
				return
			}

			return out, h.warn(ctx, args)
		},
	}
}

type (
	logErrorArgs struct {
		hasMessage bool
		Message    string

		hasFields bool
		Fields    map[string]string
	}
)

//
//
// expects implementation of error function:
// func (h log) error(ctx context.Context, args *logErrorArgs) (results *logErrorResults, err error) {
//    return
// }
func (h logHandler) Error() *atypes.Function {
	return &atypes.Function{
		Ref:  "logError",
		Type: "",
		Meta: &atypes.FunctionMeta{
			Short: "Writes error log message",
		},

		Parameters: []*atypes.Param{
			{
				Name:  "message",
				Types: []string{"String"}, Required: true,
			},
			{
				Name:  "fields",
				Types: []string{"KV"},
			},
		},

		Handler: func(ctx context.Context, in expr.Vars) (out expr.Vars, err error) {
			var (
				args = &logErrorArgs{
					hasMessage: in.Has("message"),
					hasFields:  in.Has("fields"),
				}
			)

			if err = in.Decode(args); err != nil {
				return
			}

			return out, h.error(ctx, args)
		},
	}
}
