package automation

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
// compose/automation/modules_handler.yaml

import (
	"context"
	atypes "github.com/cortezaproject/corteza-server/automation/types"
	"github.com/cortezaproject/corteza-server/compose/types"
	"github.com/cortezaproject/corteza-server/pkg/expr"
	"github.com/cortezaproject/corteza-server/pkg/logger"
	"github.com/cortezaproject/corteza-server/pkg/wfexec"
	"go.uber.org/zap"
)

var _ wfexec.ExecResponse

type (
	modulesHandlerRegistry interface {
		AddFunctions(ff ...*atypes.Function)
		Type(ref string) expr.Type
	}
)

func (h modulesHandler) register() {
	h.reg.AddFunctions(
		h.Lookup(),
	)
}

type (
	modulesLookupArgs struct {
		log *zap.Logger

		hasModule bool

		Module       interface{}
		moduleID     uint64
		moduleHandle string

		hasNamespace bool

		Namespace       interface{}
		namespaceID     uint64
		namespaceHandle string
		namespaceRes    *types.Namespace
	}

	modulesLookupResults struct {
		Module *types.Module
	}
)

//
//
// expects implementation of lookup function:
// func (h modules) lookup(ctx context.Context, args *modulesLookupArgs) (results *modulesLookupResults, err error) {
//    return
// }
func (h modulesHandler) Lookup() *atypes.Function {
	return &atypes.Function{
		Ref: "composeModulesLookup",
		Meta: &atypes.FunctionMeta{
			Short: "Lookup for compose module by ID",
		},

		Parameters: []*atypes.Param{
			{
				Name:  "module",
				Types: []string{"ID", "String"}, Required: true,
			},
			{
				Name:  "namespace",
				Types: []string{"ID", "String", "ComposeNamespace"}, Required: true,
			},
		},

		Results: []*atypes.Param{

			{
				Name:  "module",
				Types: []string{"ComposeModule"},
			},
		},

		Handler: func(ctx context.Context, in expr.Vars) (out expr.Vars, err error) {
			var (
				args = &modulesLookupArgs{
					log: logger.ContextValue(ctx, zap.NewNop()).
						With(zap.String("function", "composeModulesLookup")),
					hasModule:    in.Has("module"),
					hasNamespace: in.Has("namespace"),
				}

				results *modulesLookupResults
			)

			if err = in.Decode(args); err != nil {
				return
			}

			// Converting Module to go type
			switch casted := args.Module.(type) {
			case uint64:
				args.moduleID = casted
			case string:
				args.moduleHandle = casted
			}

			// Converting Namespace to go type
			switch casted := args.Namespace.(type) {
			case uint64:
				args.namespaceID = casted
			case string:
				args.namespaceHandle = casted
			case *types.Namespace:
				args.namespaceRes = casted
			}

			if results, err = h.lookup(ctx, args); err != nil {
				return
			}

			out = expr.Vars{}
			if out["module"], err = h.reg.Type("ComposeModule").Cast(results.Module); err != nil {
				return nil, err
			}

			return
		},
	}
}
