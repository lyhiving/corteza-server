package automation

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
// compose/automation/records_handler.yaml

import (
	"context"
	atypes "github.com/cortezaproject/corteza-server/automation/types"
	"github.com/cortezaproject/corteza-server/compose/types"
	"github.com/cortezaproject/corteza-server/pkg/expr"
	"github.com/cortezaproject/corteza-server/pkg/label"
	"github.com/cortezaproject/corteza-server/pkg/logger"
	"github.com/cortezaproject/corteza-server/pkg/wfexec"
	"go.uber.org/zap"
)

var _ wfexec.ExecResponse

type (
	recordsHandlerRegistry interface {
		AddFunctions(ff ...*atypes.Function)
		Type(ref string) expr.Type
	}
)

func (h recordsHandler) register() {
	h.reg.AddFunctions(
		h.LookupByID(),
		h.Save(),
		h.Create(),
		h.Update(),
		h.Delete(),
	)
}

type (
	recordsLookupByIDArgs struct {
		log *zap.Logger

		hasRecordID bool

		RecordID uint64

		hasModule bool

		Module       interface{}
		moduleID     uint64
		moduleHandle string
		moduleRes    *types.Module

		hasNamespace bool

		Namespace       interface{}
		namespaceID     uint64
		namespaceHandle string
		namespaceRes    *types.Namespace
	}

	recordsLookupByIDResults struct {
		Record *types.Record
	}
)

//
//
// expects implementation of lookupByID function:
// func (h records) lookupByID(ctx context.Context, args *recordsLookupByIDArgs) (results *recordsLookupByIDResults, err error) {
//    return
// }
func (h recordsHandler) LookupByID() *atypes.Function {
	return &atypes.Function{
		Ref: "composeRecordsLookupByID",
		Meta: &atypes.FunctionMeta{
			Short: "Lookup for compose record by ID",
		},

		Parameters: []*atypes.Param{
			{
				Name:  "recordID",
				Types: []string{"ID"}, Required: true,
			},
			{
				Name:  "module",
				Types: []string{"ID", "String", "ComposeModule"}, Required: true,
				Meta: &atypes.ParamMeta{
					Label:       "Module to set record type",
					Description: "Even with unique record ID across all modules, module needs to be known\nbefore doing any record operations. Mainly because records of different\nmodules can be located in different stores.",
				},
			},
			{
				Name:  "namespace",
				Types: []string{"ID", "String", "ComposeNamespace"}, Required: true,
			},
		},

		Results: []*atypes.Param{

			{
				Name:  "record",
				Types: []string{"ComposeRecord"},
			},
		},

		Handler: func(ctx context.Context, in expr.Vars) (out expr.Vars, err error) {
			var (
				args = &recordsLookupByIDArgs{
					log: logger.ContextValue(ctx, zap.NewNop()).
						With(zap.String("function", "composeRecordsLookupByID")),
					hasRecordID:  in.Has("recordID"),
					hasModule:    in.Has("module"),
					hasNamespace: in.Has("namespace"),
				}

				results *recordsLookupByIDResults
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
			case *types.Module:
				args.moduleRes = casted
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

			if results, err = h.lookupByID(ctx, args); err != nil {
				return
			}

			out = expr.Vars{}
			if out["record"], err = h.reg.Type("ComposeRecord").Cast(results.Record); err != nil {
				return nil, err
			}

			return
		},
	}
}

type (
	recordsSaveArgs struct {
		log *zap.Logger

		hasRecord bool

		Record *types.Record
	}

	recordsSaveResults struct {
		Record *types.Record
	}
)

//
//
// expects implementation of save function:
// func (h records) save(ctx context.Context, args *recordsSaveArgs) (results *recordsSaveResults, err error) {
//    return
// }
func (h recordsHandler) Save() *atypes.Function {
	return &atypes.Function{
		Ref: "composeRecordsSave",
		Meta: &atypes.FunctionMeta{
			Short: "Save record",
		},

		Parameters: []*atypes.Param{
			{
				Name:  "record",
				Types: []string{"ComposeRecord"}, Required: true,
			},
		},

		Results: []*atypes.Param{

			{
				Name:  "record",
				Types: []string{"ComposeRecord"},
			},
		},

		Handler: func(ctx context.Context, in expr.Vars) (out expr.Vars, err error) {
			var (
				args = &recordsSaveArgs{
					log: logger.ContextValue(ctx, zap.NewNop()).
						With(zap.String("function", "composeRecordsSave")),
					hasRecord: in.Has("record"),
				}

				results *recordsSaveResults
			)

			if err = in.Decode(args); err != nil {
				return
			}

			if results, err = h.save(ctx, args); err != nil {
				return
			}

			out = expr.Vars{}
			if out["record"], err = h.reg.Type("ComposeRecord").Cast(results.Record); err != nil {
				return nil, err
			}

			return
		},
	}
}

type (
	recordsCreateArgs struct {
		log *zap.Logger

		hasModule bool

		Module       interface{}
		moduleID     uint64
		moduleHandle string
		moduleRes    *types.Module

		hasNamespace bool

		Namespace       interface{}
		namespaceID     uint64
		namespaceHandle string
		namespaceRes    *types.Namespace

		hasValues bool

		Values types.RecordValueSet

		hasLabels bool

		Labels label.Labels

		hasOwnedBy bool

		OwnedBy uint64
	}

	recordsCreateResults struct {
		Record *types.Record
	}
)

//
//
// expects implementation of create function:
// func (h records) create(ctx context.Context, args *recordsCreateArgs) (results *recordsCreateResults, err error) {
//    return
// }
func (h recordsHandler) Create() *atypes.Function {
	return &atypes.Function{
		Ref: "composeRecordsCreate",
		Meta: &atypes.FunctionMeta{
			Short: "Creates and stores a new record",
		},

		Parameters: []*atypes.Param{
			{
				Name:  "module",
				Types: []string{"ID", "String", "ComposeModule"}, Required: true,
				Meta: &atypes.ParamMeta{
					Label:       "Module to set record type",
					Description: "Even with unique record ID across all modules, module needs to be known\nbefore doing any record operations. Mainly because records of different\nmodules can be located in different stores.",
				},
			},
			{
				Name:  "namespace",
				Types: []string{"ID", "String", "ComposeNamespace"}, Required: true,
			},
			{
				Name:  "values",
				Types: []string{"KV"},
			},
			{
				Name:  "labels",
				Types: []string{"KV"},
			},
			{
				Name:  "ownedBy",
				Types: []string{"ID"},
				Meta: &atypes.ParamMeta{
					Label:  "Record owner",
					Visual: map[string]interface{}{"ref": "users"},
				},
			},
		},

		Results: []*atypes.Param{

			{
				Name:  "record",
				Types: []string{"ComposeRecord"},
			},
		},

		Handler: func(ctx context.Context, in expr.Vars) (out expr.Vars, err error) {
			var (
				args = &recordsCreateArgs{
					log: logger.ContextValue(ctx, zap.NewNop()).
						With(zap.String("function", "composeRecordsCreate")),
					hasModule:    in.Has("module"),
					hasNamespace: in.Has("namespace"),
					hasValues:    in.Has("values"),
					hasLabels:    in.Has("labels"),
					hasOwnedBy:   in.Has("ownedBy"),
				}

				results *recordsCreateResults
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
			case *types.Module:
				args.moduleRes = casted
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

			if results, err = h.create(ctx, args); err != nil {
				return
			}

			out = expr.Vars{}
			if out["record"], err = h.reg.Type("ComposeRecord").Cast(results.Record); err != nil {
				return nil, err
			}

			return
		},
	}
}

type (
	recordsUpdateArgs struct {
		log *zap.Logger

		hasModule bool

		Module       interface{}
		moduleID     uint64
		moduleHandle string
		moduleRes    *types.Module

		hasNamespace bool

		Namespace       interface{}
		namespaceID     uint64
		namespaceHandle string
		namespaceRes    *types.Namespace

		hasValues bool

		Values types.RecordValueSet

		hasLabels bool

		Labels label.Labels

		hasOwnedBy bool

		OwnedBy uint64
	}

	recordsUpdateResults struct {
		Record *types.Record
	}
)

//
//
// expects implementation of update function:
// func (h records) update(ctx context.Context, args *recordsUpdateArgs) (results *recordsUpdateResults, err error) {
//    return
// }
func (h recordsHandler) Update() *atypes.Function {
	return &atypes.Function{
		Ref: "composeRecordsUpdate",
		Meta: &atypes.FunctionMeta{
			Short: "Updates an existing record",
		},

		Parameters: []*atypes.Param{
			{
				Name:  "module",
				Types: []string{"ID", "String", "ComposeModule"}, Required: true,
				Meta: &atypes.ParamMeta{
					Label:       "Module to set record type",
					Description: "Even with unique record ID across all modules, module needs to be known\nbefore doing any record operations. Mainly because records of different\nmodules can be located in different stores.",
				},
			},
			{
				Name:  "namespace",
				Types: []string{"ID", "String", "ComposeNamespace"}, Required: true,
			},
			{
				Name:  "values",
				Types: []string{"KV"},
			},
			{
				Name:  "labels",
				Types: []string{"KV"},
			},
			{
				Name:  "ownedBy",
				Types: []string{"ID"},
				Meta: &atypes.ParamMeta{
					Label:  "Record owner",
					Visual: map[string]interface{}{"ref": "users"},
				},
			},
		},

		Results: []*atypes.Param{

			{
				Name:  "record",
				Types: []string{"ComposeRecord"},
			},
		},

		Handler: func(ctx context.Context, in expr.Vars) (out expr.Vars, err error) {
			var (
				args = &recordsUpdateArgs{
					log: logger.ContextValue(ctx, zap.NewNop()).
						With(zap.String("function", "composeRecordsUpdate")),
					hasModule:    in.Has("module"),
					hasNamespace: in.Has("namespace"),
					hasValues:    in.Has("values"),
					hasLabels:    in.Has("labels"),
					hasOwnedBy:   in.Has("ownedBy"),
				}

				results *recordsUpdateResults
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
			case *types.Module:
				args.moduleRes = casted
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

			if results, err = h.update(ctx, args); err != nil {
				return
			}

			out = expr.Vars{}
			if out["record"], err = h.reg.Type("ComposeRecord").Cast(results.Record); err != nil {
				return nil, err
			}

			return
		},
	}
}

type (
	recordsDeleteArgs struct {
		log *zap.Logger

		hasRecordID bool

		RecordID uint64

		hasModule bool

		Module       interface{}
		moduleID     uint64
		moduleHandle string
		moduleRes    *types.Module

		hasNamespace bool

		Namespace       interface{}
		namespaceID     uint64
		namespaceHandle string
		namespaceRes    *types.Namespace
	}
)

//
//
// expects implementation of delete function:
// func (h records) delete(ctx context.Context, args *recordsDeleteArgs) (results *recordsDeleteResults, err error) {
//    return
// }
func (h recordsHandler) Delete() *atypes.Function {
	return &atypes.Function{
		Ref: "composeRecordsDelete",
		Meta: &atypes.FunctionMeta{
			Short: "Soft deletes compose record by ID",
		},

		Parameters: []*atypes.Param{
			{
				Name:  "recordID",
				Types: []string{"ID"}, Required: true,
			},
			{
				Name:  "module",
				Types: []string{"ID", "String", "ComposeModule"}, Required: true,
				Meta: &atypes.ParamMeta{
					Label:       "Module to set record type",
					Description: "Even with unique record ID across all modules, module needs to be known\nbefore doing any record operations. Mainly because records of different\nmodules can be located in different stores.",
				},
			},
			{
				Name:  "namespace",
				Types: []string{"ID", "String", "ComposeNamespace"}, Required: true,
			},
		},

		Handler: func(ctx context.Context, in expr.Vars) (out expr.Vars, err error) {
			var (
				args = &recordsDeleteArgs{
					log: logger.ContextValue(ctx, zap.NewNop()).
						With(zap.String("function", "composeRecordsDelete")),
					hasRecordID:  in.Has("recordID"),
					hasModule:    in.Has("module"),
					hasNamespace: in.Has("namespace"),
				}
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
			case *types.Module:
				args.moduleRes = casted
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

			return out, h.delete(ctx, args)
		},
	}
}
