package websocket

import (
	"context"
	"github.com/crusttech/crust/sam/service"
	"github.com/crusttech/crust/sam/types"
	"github.com/titpetric/factory"
	"time"
)

type (
	eventQueuePuller interface {
		EventQueuePull(origin uint64) ([]*types.EventQueueItem, error)
		EventQueueSync(origin uint64, ID uint64) error
	}
	eventQueuePusher interface {
		EventQueuePush(*types.EventQueueItem) error
	}

	eventQueueWalker interface {
		Walk(func(session *Session))
	}

	eventQueue struct {
		origin uint64
		pubsub *service.PubSub
		queue  chan *types.EventQueueItem
	}
)

const (
	eventQueueBacklog = 512
)

var eq *eventQueue

func init() {
	eq = EventQueue(factory.Sonyflake.NextID())
}

func EventQueue(origin uint64) *eventQueue {
	return &eventQueue{
		origin: origin,
		queue:  make(chan *types.EventQueueItem, eventQueueBacklog),
	}
}

func (eq *eventQueue) store(ctx context.Context, qp eventQueuePusher) {
	go func() {
		for {
			select {
			case <-ctx.Done():
			case eqi := <-eq.queue:
				qp.EventQueuePush(eqi)
			}
		}
	}()
}

func (eq *eventQueue) feedSessions(ctx context.Context, config Configuration, qp eventQueuePuller, store eventQueueWalker) error {
	newMessageEvent := make(chan struct{}, eventQueueBacklog)
	done := make(chan error, 1)

	// feed events from redis into newMessageEvent channel
	if config.pubSubMode == "redis" && config.pubSubRedis != "" {
		onConnect := func() error {
			return nil
		}
		onMessage := func(message string, payload []byte) error {
			newMessageEvent <- struct{}{}
			return nil
		}
		pubsub := service.PubSub{}.New(config.pubSubRedis, ctx)
		go func() {
			done <- pubsub.Subscribe(onConnect, onMessage, "events")
		}()
	}

	if config.pubSubMode == "poll" {
		polling := func() error {
			for {
				select {
				case <-ctx.Done():
				case <-time.After(config.pubSubInterval):
					newMessageEvent <- struct{}{}
				}
			}
		}
		go func() {
			done <- polling()
		}()
	}

	poll := func() error {
		for {
			items, err := qp.EventQueuePull(eq.origin)
			if err != nil {
				return err
			}
			if len(items) == 0 {
				return nil
			}

			var lastSyncedId uint64

			for _, item := range items {
				if item.Subscriber == "" {
					// Distribute payload to all connected sessions
					store.Walk(func(s *Session) {
						s.sendBytes(item.Payload)
					})
				} else {
					// Distribute payload to specific subscribers
					store.Walk(func(s *Session) {
						if s.subs.Get(item.Subscriber) != nil {
							s.sendBytes(item.Payload)
						}
					})
				}

				lastSyncedId = item.ID

			}

			if lastSyncedId > 0 {
				qp.EventQueueSync(eq.origin, lastSyncedId)
			}
		}
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-newMessageEvent:
			if err := poll(); err != nil {
				return err
			}
		case err := <-done:
			return err
		}
	}
}

// Adds origin to the event and puts it into queue.
func (eq *eventQueue) push(ctx context.Context, eqi *types.EventQueueItem) {
	eqi.Origin = eq.origin

	select {
	case <-ctx.Done():
	case eq.queue <- eqi:
	}
}