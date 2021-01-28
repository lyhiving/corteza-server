package automation

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
// compose/automation/namespaces_handler.yaml

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
	namespacesHandlerRegistry interface {
		AddFunctions(ff ...*atypes.Function)
		Type(ref string) expr.Type
	}
)

func (h namespacesHandler) register() {
	h.reg.AddFunctions(
		h.Lookup(),
	)
}

type (
	namespacesLookupArgs struct {
		log *zap.Logger

		hasNamespace bool

		Namespace       interface{}
		namespaceID     uint64
		namespaceHandle string
	}

	namespacesLookupResults struct {
		Namespace *types.Namespace
	}
)

//
//
// expects implementation of lookup function:
// func (h namespaces) lookup(ctx context.Context, args *namespacesLookupArgs) (results *namespacesLookupResults, err error) {
//    return
// }
func (h namespacesHandler) Lookup() *atypes.Function {
	return &atypes.Function{
		Ref: "composeNamespacesLookup",
		Meta: &atypes.FunctionMeta{
			Short: "Lookup for compose namespace by ID",
		},

		Parameters: []*atypes.Param{
			{
				Name:  "namespace",
				Types: []string{"ID", "String"}, Required: true,
			},
		},

		Results: []*atypes.Param{

			{
				Name:  "namespace",
				Types: []string{"ComposeNamespace"},
			},
		},

		Handler: func(ctx context.Context, in expr.Vars) (out expr.Vars, err error) {
			var (
				args = &namespacesLookupArgs{
					log: logger.ContextValue(ctx, zap.NewNop()).
						With(zap.String("function", "composeNamespacesLookup")),
					hasNamespace: in.Has("namespace"),
				}

				results *namespacesLookupResults
			)

			if err = in.Decode(args); err != nil {
				return
			}

			// Converting Namespace to go type
			switch casted := args.Namespace.(type) {
			case uint64:
				args.namespaceID = casted
			case string:
				args.namespaceHandle = casted
			}

			if results, err = h.lookup(ctx, args); err != nil {
				return
			}

			out = expr.Vars{}
			if out["namespace"], err = h.reg.Type("ComposeNamespace").Cast(results.Namespace); err != nil {
				return nil, err
			}

			return
		},
	}
}
