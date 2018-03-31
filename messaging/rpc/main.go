package rpc

import (
	"encoding/json"

	"github.com/ws-chart/core"
	"github.com/ws-chart/messaging/client"
	"github.com/ws-chart/protocol"
)

var messageService = newMessages()
var typingService = newTyping()

func CallMethod(c *client.Client, r *protocol.RPC) {
	processMsg := func(obj interface{}) {
		byteData, _ := json.Marshal(r.Body)
		err := json.Unmarshal(byteData, obj)
		if err != nil {
			core.Logger.Errorf("MESSAGING: Unable to read message %v %v\n", r.Body, err)
		}
	}

	switch r.Method {
	case protocol.RPC_MESSAGE_SEND:
		params := protocol.RPC_MessageSendParams{}
		processMsg(&params)
		messageService.Send(c, r.RequestId, &params)

	case protocol.RPC_MESSAGE_DELIVERED:
		params := protocol.RPC_MessageDeliveredParams{}
		processMsg(&params)
		messageService.Delivered(c, r.RequestId, &params)

	case protocol.RPC_MESSAGE_READ:
		params := protocol.RPC_MessageReadParams{}
		processMsg(&params)
		messageService.Read(c, r.RequestId, &params)

	case protocol.RPC_TYPING_START:
		params := protocol.RPC_TypingStartParams{}
		processMsg(&params)
		typingService.Start(c, r.RequestId, &params)

	case protocol.RPC_TYPING_END:
		params := protocol.RPC_TypingEndParams{}
		processMsg(&params)
		typingService.End(c, r.RequestId, &params)
	}
}
