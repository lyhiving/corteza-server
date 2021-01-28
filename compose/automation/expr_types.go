package automation

import (
	"fmt"
	"github.com/cortezaproject/corteza-server/compose/types"
	"github.com/cortezaproject/corteza-server/pkg/expr"
	"github.com/spf13/cast"
	"strings"
)

func (t *Record) Dict() map[string]interface{} {
	return t.value.Dict()
}

func (t *Record) Set(new interface{}, pp ...string) (err error) {
	if tv, is := new.(expr.TypedValue); is {
		new = tv.Get()
	}

	if len(pp) == 0 {
		var (
			ok  bool
			aux *types.Record
		)
		aux, ok = new.(*types.Record)
		if !ok {
			return fmt.Errorf("unable to cast type %T to types.Record", new)
		}

		t.value = aux
	} else {
		return setRecordProps(t.value, new, pp...)
	}
	return nil
}

func setRecordProps(rec *types.Record, val interface{}, pp ...string) (err error) {
	switch pp[0] {
	case "ID":
		return expr.SetIDWithPath(&rec.ID, val, pp[1:]...)

	case "moduleID":
		return expr.SetIDWithPath(&rec.ModuleID, val, pp[1:]...)

	case "namespaceID":
		return expr.SetIDWithPath(&rec.NamespaceID, val, pp[1:]...)

	case "values":
		return setRecordValuesWithPath(&rec.Values, val, pp[1:]...)

	case "labels":
		return expr.SetKVWithPath(&rec.Labels, val, pp[1:]...)

	case "ownedBy":
		return expr.SetIDWithPath(&rec.OwnedBy, val, pp[1:]...)

	case "createdAt":
		return expr.SetTimeWithPath(&rec.CreatedAt, val, pp[1:]...)

	case "createdBy":
		return expr.SetIDWithPath(&rec.CreatedBy, val, pp[1:]...)

	case "updatedAt":
		return expr.SetTimeWithPath(rec.UpdatedAt, val, pp[1:]...)

	case "updatedBy":
		return expr.SetIDWithPath(&rec.UpdatedBy, val, pp[1:]...)

	case "deletedAt":
		return expr.SetTimeWithPath(rec.DeletedAt, val, pp[1:]...)

	case "deletedBy":
		return expr.SetIDWithPath(&rec.DeletedBy, val, pp[1:]...)

	default:
		return fmt.Errorf("unknown record field %q", pp[0])

	}
}

func (t *RecordValues) Set(new interface{}, pp ...string) (err error) {
	return setRecordValuesWithPath(&t.value, new, pp...)
}

// setRecordValuesWithPath returns updated record value set
func setRecordValuesWithPath(rvs *types.RecordValueSet, new interface{}, pp ...string) (err error) {
	if tv, is := new.(expr.TypedValue); is {
		new = tv.Get()
	}

	if len(pp) == 0 {
		switch casted := new.(type) {
		case *types.Record:
			*rvs = casted.Values

		case types.RecordValueSet:
			*rvs = casted

		case map[string]string:
			*rvs = types.RecordValueSet{}
			for k, v := range casted {
				*rvs = rvs.Set(&types.RecordValue{Name: k, Value: v})
			}

		case map[string][]string:
			*rvs = types.RecordValueSet{}
			for k, vv := range casted {
				for i, v := range vv {
					*rvs = rvs.Set(&types.RecordValue{Name: k, Value: v, Place: uint(i)})
				}
			}

		default:
			return fmt.Errorf("unable to cast type %T to types.Record", new)
		}

		return nil
	}

	if len(pp) > 2 {
		return fmt.Errorf("invalid path for record value: %q", strings.Join(pp, "."))
	}

	rv := &types.RecordValue{Name: pp[0]}
	if rv.Value, err = cast.ToStringE(new); err != nil {
		return
	}

	if len(pp) == 2 {
		if rv.Place, err = cast.ToUintE(pp[1]); err != nil {
			return
		}
	}

	*rvs = rvs.Set(rv)

	return nil
}

func (t *Module) Set(new interface{}, pp ...string) (err error) {
	if tv, is := new.(expr.TypedValue); is {
		new = tv.Get()
	}

	// @todo implement setting via path
	var (
		ok  bool
		aux *types.Module
	)
	aux, ok = new.(*types.Module)
	if !ok {
		return fmt.Errorf("unable to cast type %T to types.Module", new)
	}
	t.value = aux
	return nil
}

func (t *Namespace) Set(new interface{}, pp ...string) (err error) {
	if tv, is := new.(expr.TypedValue); is {
		new = tv.Get()
	}

	// @todo implement setting via path
	var (
		ok  bool
		aux *types.Namespace
	)

	aux, ok = new.(*types.Namespace)
	if !ok {
		return fmt.Errorf("unable to cast type %T to types.Namespace", new)
	}
	t.value = aux
	return nil
}
