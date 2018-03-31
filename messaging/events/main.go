package events

import (
	"github.com/ievgen-ma/ws-chart/messaging/client"
	"github.com/ievgen-ma/ws-chart/protocol"
)

type Event struct {
	*protocol.Event
}

func NewEvent(t protocol.Event_Type) *Event {
	pe := &protocol.Event{
		Type: t,
	}

	return &Event{pe}
}

func (e *Event) SendToClient(clientID string) {
	c := client.NewClient(clientID)
	c.SendEvent(e.Event)
}
