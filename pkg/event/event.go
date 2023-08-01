package event

type Event interface {
	Type() string
}

type ScopeEvent interface {
	GetUserID() string
}

type Subscribe struct {
	ch     chan Event
	topics []string
}

var (
	ch         chan Event
	subscribes []*Subscribe
)

func init() {
	ch = make(chan Event, 50)
	subscribes = make([]*Subscribe, 0)

	go dispatcher()
}

func dispatcher() {
	for {
		select {
		case evt := <-ch:
			dispatch(evt)
		}
	}
}

func dispatch(evt Event) {
	for _, sub := range subscribes {
		if len(sub.topics) == 0 {
			sub.ch <- evt
			continue
		}

		for _, topic := range sub.topics {
			if topic == evt.Type() {
				sub.ch <- evt
				break
			}
		}
	}
}

func Pub(evt Event) {
	ch <- evt
}

func Sub(topics ...string) chan Event {
	subscribe := &Subscribe{
		ch:     make(chan Event, 50),
		topics: topics,
	}

	subscribes = append(subscribes, subscribe)
	return subscribe.ch
}
