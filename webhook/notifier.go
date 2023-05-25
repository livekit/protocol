package webhook

import (
	"context"
	"sync"

	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/logger"
)

type QueuedNotifier interface {
	QueueNotify(ctx context.Context, event *livekit.WebhookEvent) error
}

type DefaultNotifier struct {
	urlNotifiers []*URLNotifier
}

func NewDefaultNotifier(apiKey, apiSecret string, urls []string) QueuedNotifier {
	n := &DefaultNotifier{}
	for _, url := range urls {
		u := NewURLNotifier(URLNotifierParams{
			URL:       url,
			Logger:    logger.GetLogger(),
			APIKey:    apiKey,
			APISecret: apiSecret,
		})
		n.urlNotifiers = append(n.urlNotifiers, u)
	}
	return n
}

func (n *DefaultNotifier) Stop(force bool) {
	wg := sync.WaitGroup{}
	for _, u := range n.urlNotifiers {
		wg.Add(1)
		go func(u *URLNotifier) {
			defer wg.Done()
			u.Stop(force)
		}(u)
	}
	wg.Wait()
}

func (n *DefaultNotifier) QueueNotify(_ context.Context, event *livekit.WebhookEvent) error {
	for _, u := range n.urlNotifiers {
		if err := u.QueueNotify(event); err != nil {
			return err
		}
	}
	return nil
}
