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
	"github.com/cortezaproject/corteza-server/pkg/wfexec"
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
		h.Validate(),
		h.Convert(),
		h.Create(),
		h.Update(),
		h.Delete(),
	)
}

type (
	recordsLookupByIDArgs struct {
		hasRecordID bool
		RecordID    uint64

		hasModule    bool
		Module       interface{}
		moduleID     uint64
		moduleHandle string
		moduleRes    *types.Module

		hasNamespace    bool
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
		Ref:  "composeRecordsLookupByID",
		Type: "",
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
				Types: []string{"ID", "Handle", "ComposeModule"}, Required: true,
				Meta: &atypes.ParamMeta{
					Label:       "Module to set record type",
					Description: "Even with unique record ID across all modules, module needs to be known\nbefore doing any record operations. Mainly because records of different\nmodules can be located in different stores.",
				},
			},
			{
				Name:  "namespace",
				Types: []string{"ID", "Handle", "ComposeNamespace"}, Required: true,
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

			var results *recordsLookupByIDResults
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
		hasRecord bool
		Record    *types.Record
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
		Ref:  "composeRecordsSave",
		Type: "",
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
					hasRecord: in.Has("record"),
				}
			)

			if err = in.Decode(args); err != nil {
				return
			}

			var results *recordsSaveResults
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
	recordsValidateArgs struct {
		hasRecord bool
		Record    *types.Record

		hasModule    bool
		Module       interface{}
		moduleID     uint64
		moduleHandle string
		moduleRes    *types.Module

		hasNamespace    bool
		Namespace       interface{}
		namespaceID     uint64
		namespaceHandle string
		namespaceRes    *types.Namespace
	}

	recordsValidateResults struct {
		Valid  bool
		Errors *types.RecordValueErrorSet
	}
)

//
//
// expects implementation of validate function:
// func (h records) validate(ctx context.Context, args *recordsValidateArgs) (results *recordsValidateResults, err error) {
//    return
// }
func (h recordsHandler) Validate() *atypes.Function {
	return &atypes.Function{
		Ref:  "composeRecordsValidate",
		Type: "",
		Meta: &atypes.FunctionMeta{
			Short: "Validate record",
		},

		Parameters: []*atypes.Param{
			{
				Name:  "record",
				Types: []string{"ComposeRecord"}, Required: true,
			},
			{
				Name:  "module",
				Types: []string{"ID", "Handle", "ComposeModule"},
				Meta: &atypes.ParamMeta{
					Label:       "Module to set record type",
					Description: "Even with unique record ID across all modules, module needs to be known\nbefore doing any record operations. Mainly because records of different\nmodules can be located in different stores.",
				},
			},
			{
				Name:  "namespace",
				Types: []string{"ID", "Handle", "ComposeNamespace"},
			},
		},

		Results: []*atypes.Param{

			{
				Name:  "valid",
				Types: []string{"Boolean"},
				Meta: &atypes.ParamMeta{
					Label: "Set to true when record is valid",
				},
			},

			{
				Name:  "errors",
				Types: []string{"RecordValueErrorSet,"},
				Meta: &atypes.ParamMeta{
					Label: "List of errors collected when validating the record",
				},
			},
		},

		Handler: func(ctx context.Context, in expr.Vars) (out expr.Vars, err error) {
			var (
				args = &recordsValidateArgs{
					hasRecord:    in.Has("record"),
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

			var results *recordsValidateResults
			if results, err = h.validate(ctx, args); err != nil {
				return
			}

			out = expr.Vars{}
			if out["valid"], err = h.reg.Type("Boolean").Cast(results.Valid); err != nil {
				return nil, err
			}
			if out["errors"], err = h.reg.Type("RecordValueErrorSet,").Cast(results.Errors); err != nil {
				return nil, err
			}

			return
		},
	}
}

type (
	recordsConvertArgs struct {
		hasSource bool
		Source    *types.Record

		hasModule    bool
		Module       interface{}
		moduleID     uint64
		moduleHandle string
		moduleRes    *types.Module

		hasNamespace    bool
		Namespace       interface{}
		namespaceID     uint64
		namespaceHandle string
		namespaceRes    *types.Namespace

		hasMap bool
		Map    map[string]string
	}

	recordsConvertResults struct {
		Record *types.Record
	}
)

//
//
// expects implementation of convert function:
// func (h records) convert(ctx context.Context, args *recordsConvertArgs) (results *recordsConvertResults, err error) {
//    return
// }
func (h recordsHandler) Convert() *atypes.Function {
	return &atypes.Function{
		Ref:  "composeRecordsConvert",
		Type: "",
		Meta: &atypes.FunctionMeta{
			Short: "Converts record",
		},

		Parameters: []*atypes.Param{
			{
				Name:  "source",
				Types: []string{"ComposeRecord"}, Required: true,
			},
			{
				Name:  "module",
				Types: []string{"ID", "Handle", "ComposeModule"}, Required: true,
				Meta: &atypes.ParamMeta{
					Label:       "Module to set record type",
					Description: "Even with unique record ID across all modules, module needs to be known\nbefore doing any record operations. Mainly because records of different\nmodules can be located in different stores.",
				},
			},
			{
				Name:  "namespace",
				Types: []string{"ID", "Handle", "ComposeNamespace"}, Required: true,
			},
			{
				Name:  "map",
				Types: []string{"KV"},
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
				args = &recordsConvertArgs{
					hasSource:    in.Has("source"),
					hasModule:    in.Has("module"),
					hasNamespace: in.Has("namespace"),
					hasMap:       in.Has("map"),
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

			var results *recordsConvertResults
			if results, err = h.convert(ctx, args); err != nil {
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
		hasModule    bool
		Module       interface{}
		moduleID     uint64
		moduleHandle string
		moduleRes    *types.Module

		hasNamespace    bool
		Namespace       interface{}
		namespaceID     uint64
		namespaceHandle string
		namespaceRes    *types.Namespace

		hasValues bool
		Values    types.RecordValueSet

		hasLabels bool
		Labels    label.Labels

		hasOwnedBy bool
		OwnedBy    uint64
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
		Ref:  "composeRecordsCreate",
		Type: "",
		Meta: &atypes.FunctionMeta{
			Short: "Creates and stores a new record",
		},

		Parameters: []*atypes.Param{
			{
				Name:  "module",
				Types: []string{"ID", "Handle", "ComposeModule"}, Required: true,
				Meta: &atypes.ParamMeta{
					Label:       "Module to set record type",
					Description: "Even with unique record ID across all modules, module needs to be known\nbefore doing any record operations. Mainly because records of different\nmodules can be located in different stores.",
				},
			},
			{
				Name:  "namespace",
				Types: []string{"ID", "Handle", "ComposeNamespace"}, Required: true,
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
					hasModule:    in.Has("module"),
					hasNamespace: in.Has("namespace"),
					hasValues:    in.Has("values"),
					hasLabels:    in.Has("labels"),
					hasOwnedBy:   in.Has("ownedBy"),
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

			var results *recordsCreateResults
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
		hasModule    bool
		Module       interface{}
		moduleID     uint64
		moduleHandle string
		moduleRes    *types.Module

		hasNamespace    bool
		Namespace       interface{}
		namespaceID     uint64
		namespaceHandle string
		namespaceRes    *types.Namespace

		hasValues bool
		Values    types.RecordValueSet

		hasLabels bool
		Labels    label.Labels

		hasOwnedBy bool
		OwnedBy    uint64
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
		Ref:  "composeRecordsUpdate",
		Type: "",
		Meta: &atypes.FunctionMeta{
			Short: "Updates an existing record",
		},

		Parameters: []*atypes.Param{
			{
				Name:  "module",
				Types: []string{"ID", "Handle", "ComposeModule"}, Required: true,
				Meta: &atypes.ParamMeta{
					Label:       "Module to set record type",
					Description: "Even with unique record ID across all modules, module needs to be known\nbefore doing any record operations. Mainly because records of different\nmodules can be located in different stores.",
				},
			},
			{
				Name:  "namespace",
				Types: []string{"ID", "Handle", "ComposeNamespace"}, Required: true,
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
					hasModule:    in.Has("module"),
					hasNamespace: in.Has("namespace"),
					hasValues:    in.Has("values"),
					hasLabels:    in.Has("labels"),
					hasOwnedBy:   in.Has("ownedBy"),
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

			var results *recordsUpdateResults
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
		hasRecordID bool
		RecordID    uint64

		hasModule    bool
		Module       interface{}
		moduleID     uint64
		moduleHandle string
		moduleRes    *types.Module

		hasNamespace    bool
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
		Ref:  "composeRecordsDelete",
		Type: "",
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
				Types: []string{"ID", "Handle", "ComposeModule"}, Required: true,
				Meta: &atypes.ParamMeta{
					Label:       "Module to set record type",
					Description: "Even with unique record ID across all modules, module needs to be known\nbefore doing any record operations. Mainly because records of different\nmodules can be located in different stores.",
				},
			},
			{
				Name:  "namespace",
				Types: []string{"ID", "Handle", "ComposeNamespace"}, Required: true,
			},
		},

		Handler: func(ctx context.Context, in expr.Vars) (out expr.Vars, err error) {
			var (
				args = &recordsDeleteArgs{
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
