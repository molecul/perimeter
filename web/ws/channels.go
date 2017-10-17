package ws

import "github.com/dustin/go-broadcast"

var channels = make(map[string]broadcast.Broadcaster)

func OpenListener(channel string) chan interface{} {
	listener := make(chan interface{})

	GetChannel(channel).Register(listener)

	return listener
}

func CloseListener(channel string, listener chan interface{}) {
	GetChannel(channel).Unregister(listener)

	close(listener)
}

func GetChannel(channel string) broadcast.Broadcaster {
	b, ok := channels[channel]

	if !ok {
		b = broadcast.NewBroadcaster(10)
		channels[channel] = b
	}

	return b
}
