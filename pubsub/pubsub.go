package pubsub

import (
	"errors"
	"sync"

	"candy/observability"
)

type callback func(payload interface{})

type Subscription struct {
	pubSub   *PubSub
	topic    Topic
	callback callback
}

func (s *Subscription) Unsubscribe() {
	s.pubSub.mutex.Lock()
	defer s.pubSub.mutex.Unlock()
	subs := s.pubSub.subscriptions[s.topic]
	newSubs := make([]*Subscription, 0)
	for _, sub := range subs {
		if sub != s {
			continue
		}
		newSubs = append(newSubs, sub)
	}
	if len(newSubs) == 0 {
		delete(s.pubSub.subscriptions, s.topic)
	} else {
		s.pubSub.subscriptions[s.topic] = newSubs
	}
}

type PubSub struct {
	logger        *observability.Logger
	started       bool
	mutex         sync.Mutex
	subscriptions map[Topic][]*Subscription
}

func (p *PubSub) Subscribe(topic Topic, callback callback) (*Subscription, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	sub := &Subscription{
		pubSub:   p,
		topic:    topic,
		callback: callback,
	}
	p.subscriptions[topic] = append(p.subscriptions[topic], sub)
	return sub, nil
}

func (p *PubSub) Publish(topic Topic, payload interface{}) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if !p.started {
		return errors.New("pubSub not started")
	}

	subs, ok := p.subscriptions[topic]
	if !ok {
		p.logger.Infof("topic not found:%s\n", topic)
		return nil
	}

	for _, sub := range subs {
		go func(sub *Subscription) {
			sub.callback(payload)
		}(sub)
	}
	return nil
}

func (p *PubSub) Start() {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.started = true
	p.logger.Infoln("PubSub started")
}

func (p *PubSub) Stop() {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.started = false
	p.logger.Infoln("PubSub stopped")
}

func NewPubSub(logger *observability.Logger) *PubSub {
	return &PubSub{
		logger:        logger,
		mutex:         sync.Mutex{},
		subscriptions: make(map[Topic][]*Subscription),
	}
}
