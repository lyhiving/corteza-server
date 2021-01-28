package automation

import (
	"context"
	"github.com/cortezaproject/corteza-server/compose/types"
)

type (
	namespaceService interface {
		FindByID(ctx context.Context, namespaceID uint64) (*types.Namespace, error)
		FindByHandle(ctx context.Context, handle string) (*types.Namespace, error)
		Find(ctx context.Context, filter types.NamespaceFilter) (set types.NamespaceSet, f types.NamespaceFilter, err error)

		Create(ctx context.Context, namespace *types.Namespace) (*types.Namespace, error)
		Update(ctx context.Context, namespace *types.Namespace) (*types.Namespace, error)

		DeleteByID(ctx context.Context, namespaceID uint64) error
	}

	namespacesHandler struct {
		reg namespacesHandlerRegistry
		ns  namespaceService
	}
)

func NamespacesHandler(reg namespacesHandlerRegistry, ns namespaceService) *namespacesHandler {
	h := &namespacesHandler{
		reg: reg,
		ns:  ns,
	}

	h.register()
	return h
}

func (h namespacesHandler) lookup(ctx context.Context, args *namespacesLookupArgs) (results *namespacesLookupResults, err error) {
	results = &namespacesLookupResults{}

	if args.namespaceID > 0 {
		results.Namespace, err = h.ns.FindByID(ctx, args.namespaceID)
	} else {
		results.Namespace, err = h.ns.FindByHandle(ctx, args.namespaceHandle)
	}

	return
}
