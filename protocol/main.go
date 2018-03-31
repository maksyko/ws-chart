package protocol

type RPC_Method int

const (
	RPC_MESSAGE_SEND      RPC_Method = 40
	RPC_MESSAGE_DELIVERED RPC_Method = 41
	RPC_MESSAGE_READ      RPC_Method = 42
	RPC_TYPING_START      RPC_Method = 60
	RPC_TYPING_END        RPC_Method = 61
)

type RPC struct {
	Method    RPC_Method  `json:"method"`
	RequestId string      `json:"request_id"`
	Body      interface{} `json:"body,omitempty"`
}

type RPC_MessageSendParams struct {
	Data string `json:"data"`
}

type RPC_MessageDeliveredParams struct {
	MessageId string `json:"message_id,omitempty"`
}
type RPC_MessageReadParams struct {
	MessageId string `json:"message_id,omitempty"`
}
type RPC_TypingStartParams struct{}
type RPC_TypingEndParams struct{}

// =====================================================================================================================

type Event_Type int

const (
	Event_MESSAGE           Event_Type = 20
	Event_MESSAGE_SENT      Event_Type = 21
	Event_MESSAGE_DELIVERED Event_Type = 22
	Event_MESSAGE_READ      Event_Type = 23
	Event_TYPING_START      Event_Type = 40
	Event_TYPING_END        Event_Type = 41
)

type Event struct {
	Type       Event_Type  `json:"type"`
	ResponseTo string      `json:"response_to"`
	Body       interface{} `json:"body,omitempty"`
}

type Event_Message struct {
	MessageId string `json:"message_id,omitempty"`
	Data      string `json:"data,omitempty"`
}

type Event_MessageSent struct {
	MessageId string `json:"message_id,omitempty"`
}

type Event_MessageDelivered struct {
	MessageId string `json:"message_id,omitempty"`
}

type Event_MessageRead struct {
	MessageId string `json:"message_id,omitempty"`
}
