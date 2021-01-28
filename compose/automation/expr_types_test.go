package automation

import (
	"github.com/cortezaproject/corteza-server/compose/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSetRecordValuesWithPath(t *testing.T) {
	var (
		r   = require.New(t)
		rvs = types.RecordValueSet{}
	)

	r.NoError(setRecordValuesWithPath(&rvs, "a", "field1"))
	r.NoError(setRecordValuesWithPath(&rvs, "a", "field1", "1"))
	r.True(rvs.Has("field1", 0))
	r.True(rvs.Has("field1", 1))

	rvs = types.RecordValueSet{}
	r.NoError(setRecordValuesWithPath(&rvs, map[string]string{"field2": "b"}))
	r.True(rvs.Has("field2", 0))

	rvs = types.RecordValueSet{}
	r.NoError(setRecordValuesWithPath(&rvs, map[string][]string{"field2": []string{"a", "b"}}))
	r.True(rvs.Has("field2", 0))
	r.True(rvs.Has("field2", 1))
}
