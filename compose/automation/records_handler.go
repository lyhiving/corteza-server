package automation

import (
	"context"
	"fmt"
	"github.com/cortezaproject/corteza-server/compose/types"
)

type (
	recordService interface {
		FindByID(ctx context.Context, namespaceID, moduleID, recordID uint64) (*types.Record, error)
		Find(ctx context.Context, filter types.RecordFilter) (set types.RecordSet, f types.RecordFilter, err error)

		Create(ctx context.Context, record *types.Record) (*types.Record, error)
		Update(ctx context.Context, record *types.Record) (*types.Record, error)
		Bulk(ctx context.Context, oo ...*types.RecordBulkOperation) (types.RecordSet, error)

		Validate(ctx context.Context, rec *types.Record) error

		DeleteByID(ctx context.Context, namespaceID, moduleID uint64, recordID ...uint64) error
	}

	recordModuleService interface {
		FindByID(ctx context.Context, namespaceID, moduleID uint64) (*types.Module, error)
		FindByHandle(ctx context.Context, namespaceID uint64, handle string) (*types.Module, error)
	}

	recordNamespaceService interface {
		FindByID(ctx context.Context, namespaceID uint64) (*types.Namespace, error)
		FindByHandle(ctx context.Context, handle string) (*types.Namespace, error)
	}

	recordsHandler struct {
		reg recordsHandlerRegistry
		ns  recordNamespaceService
		mod recordModuleService
		rec recordService
	}
)

func RecordsHandler(reg recordsHandlerRegistry, ns recordNamespaceService, mod recordModuleService, rec recordService) *recordsHandler {
	h := &recordsHandler{
		reg: reg,
		ns:  ns,
		mod: mod,
		rec: rec,
	}

	h.register()
	return h
}

func (h recordsHandler) resolveNamespace(ctx context.Context, id *uint64, handle string, res *types.Namespace) (err error) {
	if *id == 0 {
		if len(handle) > 0 {
			if res, err = h.ns.FindByHandle(ctx, handle); err != nil {
				return
			}
		}

		if res != nil {
			*id = res.ID
		}
	}

	return
}

func (h recordsHandler) resolveModule(ctx context.Context, namespaceID uint64, id *uint64, handle string, res *types.Module) (err error) {
	if *id == 0 {
		if len(handle) > 0 {
			if res, err = h.mod.FindByHandle(ctx, namespaceID, handle); err != nil {
				return
			}
		}

		if res != nil {
			*id = res.ID
		}
	}

	return
}

func (h recordsHandler) lookupByID(ctx context.Context, args *recordsLookupByIDArgs) (results *recordsLookupByIDResults, err error) {
	results = &recordsLookupByIDResults{}

	if err = h.resolveNamespace(ctx, &args.namespaceID, args.namespaceHandle, args.namespaceRes); err != nil {
		return nil, err
	}

	if err = h.resolveModule(ctx, args.namespaceID, &args.moduleID, args.moduleHandle, args.moduleRes); err != nil {
		return nil, err
	}

	results.Record, err = h.rec.FindByID(ctx, args.namespaceID, args.moduleID, args.RecordID)
	return
}

func (h recordsHandler) create(ctx context.Context, args *recordsCreateArgs) (results *recordsCreateResults, err error) {
	results = &recordsCreateResults{}

	if err = h.resolveNamespace(ctx, &args.namespaceID, args.namespaceHandle, args.namespaceRes); err != nil {
		return nil, err
	}

	if err = h.resolveModule(ctx, args.namespaceID, &args.moduleID, args.moduleHandle, args.moduleRes); err != nil {
		return nil, err
	}

	rec := &types.Record{
		ModuleID:    args.moduleID,
		NamespaceID: args.namespaceID,
		Values:      args.Values,
		Labels:      args.Labels,
		OwnedBy:     args.OwnedBy,
	}

	results.Record, err = h.rec.Create(ctx, rec)
	return
}

func (h recordsHandler) save(ctx context.Context, args *recordsSaveArgs) (results *recordsSaveResults, err error) {
	results = &recordsSaveResults{}
	results.Record, err = h.rec.Update(ctx, args.Record)
	return
}

func (h recordsHandler) validate(ctx context.Context, args *recordsValidateArgs) (*recordsValidateResults, error) {
	if err := h.resolveNamespace(ctx, &args.namespaceID, args.namespaceHandle, args.namespaceRes); err != nil {
		return nil, err
	}

	if err := h.resolveModule(ctx, args.namespaceID, &args.moduleID, args.moduleHandle, args.moduleRes); err != nil {
		return nil, err
	}

	results := &recordsValidateResults{Valid: true, Errors: &types.RecordValueErrorSet{}}
	if err := h.rec.Validate(ctx, args.Record); err != nil {
		results.Valid = false

		if rves, is := err.(*types.RecordValueErrorSet); is {
			results.Errors = rves
		} else {
			return nil, err
		}
	}

	return results, nil
}

func (h recordsHandler) convert(ctx context.Context, args *recordsConvertArgs) (results *recordsConvertResults, err error) {
	var (
		sModule, tModule *types.Module
	)

	if sModule, err = h.mod.FindByID(ctx, args.Source.NamespaceID, args.Source.ModuleID); err != nil {
		return
	}

	if err = h.resolveNamespace(ctx, &args.namespaceID, args.namespaceHandle, args.namespaceRes); err != nil {
		return
	}

	if err = h.resolveModule(ctx, args.namespaceID, &args.moduleID, args.moduleHandle, args.moduleRes); err != nil {
		return
	}

	if args.moduleRes == nil {
		if tModule, err = h.mod.FindByID(ctx, args.namespaceID, args.moduleID); err != nil {
			return
		}
	}

	rvs := types.RecordValueSet{}

	for sField, tField := range args.Map {
		if !sModule.Fields.HasName(sField) {
			return nil, fmt.Errorf("no such field %s on target module %s", sField, sModule.Handle)
		}

		if !tModule.Fields.HasName(tField) {
			return nil, fmt.Errorf("no such field %s on target module %s", tField, tModule.Handle)
		}

		_ = args.Source.Values.FilterByName(sField).Walk(func(v *types.RecordValue) error {
			rvs = rvs.Set(v)
			return nil
		})
	}

	results = &recordsConvertResults{
		Record: &types.Record{
			ModuleID:    tModule.ID,
			NamespaceID: tModule.NamespaceID,
			Values:      rvs,
		},
	}

	return results, nil
}

func (h recordsHandler) update(ctx context.Context, args *recordsUpdateArgs) (results *recordsUpdateResults, err error) {
	results = &recordsUpdateResults{}
	if err = h.resolveNamespace(ctx, &args.namespaceID, args.namespaceHandle, args.namespaceRes); err != nil {
		return nil, err
	}

	if err = h.resolveModule(ctx, args.namespaceID, &args.moduleID, args.moduleHandle, args.moduleRes); err != nil {
		return nil, err
	}

	rec := &types.Record{
		ModuleID:    args.moduleID,
		NamespaceID: args.namespaceID,
		Values:      args.Values,
		Labels:      args.Labels,
		OwnedBy:     args.OwnedBy,
	}

	results.Record, err = h.rec.Update(ctx, rec)
	return
}

func (h recordsHandler) delete(ctx context.Context, args *recordsDeleteArgs) (err error) {
	if err = h.resolveNamespace(ctx, &args.namespaceID, args.namespaceHandle, args.namespaceRes); err != nil {
		return err
	}

	if err = h.resolveModule(ctx, args.namespaceID, &args.moduleID, args.moduleHandle, args.moduleRes); err != nil {
		return err
	}

	return h.rec.DeleteByID(ctx, args.namespaceID, args.moduleID, args.RecordID)
}
