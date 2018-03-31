package rpc

import (
	"github.com/twinj/uuid"
	"github.com/ievgen-ma/ws-chart/messaging/client"
	"github.com/ievgen-ma/ws-chart/messaging/events"
	"github.com/ievgen-ma/ws-chart/protocol"
)

type messages struct{}

func newMessages() *messages {
	return &messages{}
}

func (m *messages) Send(c *client.Client, requestID string, p *protocol.RPC_MessageSendParams) {
	messageId := uuid.NewV4().String()

	e := events.NewEvent(protocol.Event_MESSAGE)
	e.Body = protocol.Event_Message{
		MessageId: messageId,
		Data:      p.Data,
	}
	e.Event.ResponseTo = c.ID // user__1 delivered, read, typing start|end
	e.SendToClient(requestID) // user__2

	es := events.NewEvent(protocol.Event_MESSAGE_SENT)
	es.Body = protocol.Event_MessageSent{
		MessageId: messageId,
	}
	es.Event.ResponseTo = c.ID // user__1 delivered, read, typing start|end
	es.SendToClient(requestID) // user__2
}

func (m *messages) Delivered(c *client.Client, requestID string, p *protocol.RPC_MessageDeliveredParams) {
	e := events.NewEvent(protocol.Event_MESSAGE_DELIVERED)
	e.Body = protocol.Event_MessageDelivered{
		MessageId: p.MessageId,
	}
	e.Event.ResponseTo = c.ID
	e.SendToClient(requestID)
}

func (m *messages) Read(c *client.Client, requestID string, p *protocol.RPC_MessageReadParams) {
	e := events.NewEvent(protocol.Event_MESSAGE_READ)
	e.Body = protocol.Event_MessageRead{
		MessageId: p.MessageId,
	}
	e.Event.ResponseTo = c.ID
	e.SendToClient(requestID)
}
