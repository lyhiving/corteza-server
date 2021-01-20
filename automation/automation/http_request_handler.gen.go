package automation

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
// automation/automation/http_request_handler.yaml

import (
	"context"
	atypes "github.com/cortezaproject/corteza-server/automation/types"
	"github.com/cortezaproject/corteza-server/pkg/expr"
	"github.com/cortezaproject/corteza-server/pkg/wfexec"
	"io"
	"net/http"
	"net/url"
	"time"
)

var _ wfexec.ExecResponse

type (
	httpRequestHandlerRegistry interface {
		AddFunctions(ff ...*atypes.Function)
		Type(ref string) expr.Type
	}
)

func (h httpRequestHandler) register() {
	h.reg.AddFunctions(
		h.Send(),
	)
}

type (
	httpRequestSendArgs struct {
		hasUrl bool
		Url    string

		hasMethod bool
		Method    string

		hasParams bool
		Params    url.Values

		hasHeaders bool
		Headers    http.Header

		hasHeaderAuthBearer bool
		HeaderAuthBearer    string

		hasHeaderAuthUsername bool
		HeaderAuthUsername    string

		hasHeaderAuthPassword bool
		HeaderAuthPassword    string

		hasHeaderUserAgent bool
		HeaderUserAgent    string

		hasHeaderContentType bool
		HeaderContentType    string

		hasTimeout bool
		Timeout    time.Duration

		hasForm bool
		Form    url.Values

		hasBody    bool
		Body       interface{}
		bodyString string
		bodyStream io.Reader
		bodyRaw    interface{}
	}

	httpRequestSendResults struct {
		Status        string
		StatusCode    int
		Headers       http.Header
		ContentLength int64
		ContentType   string
		Body          io.Reader
	}
)

//
//
// expects implementation of send function:
// func (h httpRequest) send(ctx context.Context, args *httpRequestSendArgs) (results *httpRequestSendResults, err error) {
//    return
// }
func (h httpRequestHandler) Send() *atypes.Function {
	return &atypes.Function{
		Ref:  "httpRequestSend",
		Type: "",
		Meta: &atypes.FunctionMeta{
			Short: "Sends HTTP request",
		},

		Parameters: []*atypes.Param{
			{
				Name:  "url",
				Types: []string{"String"}, Required: true,
			},
			{
				Name:  "method",
				Types: []string{"String"}, Required: true,
			},
			{
				Name:  "params",
				Types: []string{"KVV"},
			},
			{
				Name:  "headers",
				Types: []string{"KVV"},
			},
			{
				Name:  "headerAuthBearer",
				Types: []string{"String"},
			},
			{
				Name:  "headerAuthUsername",
				Types: []string{"String"},
			},
			{
				Name:  "headerAuthPassword",
				Types: []string{"String"},
			},
			{
				Name:  "headerUserAgent",
				Types: []string{"String"},
			},
			{
				Name:  "headerContentType",
				Types: []string{"String"},
			},
			{
				Name:  "timeout",
				Types: []string{"Duration"},
			},
			{
				Name:  "form",
				Types: []string{"KVV"},
			},
			{
				Name:  "body",
				Types: []string{"String", "Reader", "Any"},
			},
		},

		Results: []*atypes.Param{

			{
				Name:  "status",
				Types: []string{"String"},
			},

			{
				Name:  "statusCode",
				Types: []string{"Integer"},
			},

			{
				Name:  "headers",
				Types: []string{"KVV"},
			},

			{
				Name:  "contentLength",
				Types: []string{"Integer"},
			},

			{
				Name:  "contentType",
				Types: []string{"String"},
			},

			{
				Name:  "body",
				Types: []string{"Reader"},
			},
		},

		Handler: func(ctx context.Context, in expr.Vars) (out expr.Vars, err error) {
			var (
				args = &httpRequestSendArgs{
					hasUrl:                in.Has("url"),
					hasMethod:             in.Has("method"),
					hasParams:             in.Has("params"),
					hasHeaders:            in.Has("headers"),
					hasHeaderAuthBearer:   in.Has("headerAuthBearer"),
					hasHeaderAuthUsername: in.Has("headerAuthUsername"),
					hasHeaderAuthPassword: in.Has("headerAuthPassword"),
					hasHeaderUserAgent:    in.Has("headerUserAgent"),
					hasHeaderContentType:  in.Has("headerContentType"),
					hasTimeout:            in.Has("timeout"),
					hasForm:               in.Has("form"),
					hasBody:               in.Has("body"),
				}
			)

			if err = in.Decode(args); err != nil {
				return
			}

			// Converting Body to go type
			switch casted := args.Body.(type) {
			case string:
				args.bodyString = casted
			case io.Reader:
				args.bodyStream = casted
			case interface{}:
				args.bodyRaw = casted
			}

			var results *httpRequestSendResults
			if results, err = h.send(ctx, args); err != nil {
				return
			}

			out = expr.Vars{}
			if out["status"], err = h.reg.Type("String").Cast(results.Status); err != nil {
				return nil, err
			}
			if out["statusCode"], err = h.reg.Type("Integer").Cast(results.StatusCode); err != nil {
				return nil, err
			}
			if out["headers"], err = h.reg.Type("KVV").Cast(results.Headers); err != nil {
				return nil, err
			}
			if out["contentLength"], err = h.reg.Type("Integer").Cast(results.ContentLength); err != nil {
				return nil, err
			}
			if out["contentType"], err = h.reg.Type("String").Cast(results.ContentType); err != nil {
				return nil, err
			}
			if out["body"], err = h.reg.Type("Reader").Cast(results.Body); err != nil {
				return nil, err
			}

			return
		},
	}
}
