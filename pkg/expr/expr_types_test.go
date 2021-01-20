package expr

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"net/url"
	"testing"
)

func TestKV_Set(t *testing.T) {
	var (
		req = require.New(t)

		vars = &KV{value: map[string]string{
			"k1": "v1",
			"k2": "v2",
		}}
	)

	req.NoError(vars.Set("v11", "k1"))
	req.Equal("v11", vars.value["k1"])

}

func TestKVV_Set(t *testing.T) {
	var (
		req = require.New(t)
		kvv KVV
	)

	req.NoError(kvv.Set("bar", "foo"))
	req.Contains(kvv.value, "foo")
	req.Equal([]string{"bar"}, kvv.value["foo"])

	// Making sure http.Header is properly converted
	kvv = KVV{}
	req.NoError(kvv.Set(http.Header{"foo": []string{"bar"}}))
	req.Contains(kvv.value, "foo")
	req.Equal([]string{"bar"}, kvv.value["foo"])

	// Making sure url.Values are properly converted
	kvv = KVV{}
	req.NoError(kvv.Set(url.Values{"foo": []string{"bar"}}))
	req.Contains(kvv.value, "foo")
	req.Equal([]string{"bar"}, kvv.value["foo"])
}
