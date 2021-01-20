package automation

import (
	"context"
	"go.uber.org/zap"
)

type (
	logHandler struct {
		reg logHandlerRegistry
	}
)

func zapFields(fields map[string]string) []zap.Field {
	ff := make([]zap.Field, 0, len(fields))
	for k, v := range fields {
		ff = append(ff, zap.String(k, v))
	}

	return ff
}

func LogHandler(reg logHandlerRegistry) *logHandler {
	h := &logHandler{
		reg: reg,
	}

	h.register()
	return h
}

func (h logHandler) debug(_ context.Context, args *logDebugArgs) (err error) {
	args.log.Debug(args.Message, zapFields(args.Fields)...)
	return nil
}

func (h logHandler) info(_ context.Context, args *logInfoArgs) (err error) {
	args.log.Info(args.Message, zapFields(args.Fields)...)
	return nil
}

func (h logHandler) warn(_ context.Context, args *logWarnArgs) (err error) {
	args.log.Warn(args.Message, zapFields(args.Fields)...)
	return nil
}

func (h logHandler) error(_ context.Context, args *logErrorArgs) (err error) {
	args.log.Error(args.Message, zapFields(args.Fields)...)
	return nil
}
