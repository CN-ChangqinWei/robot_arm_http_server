package hook

import (
	"github.com/bxcodec/go-clean-arch/forward"
	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
)

type ForwardHandler struct {
	mqtt.HookBase
	Server  *mqtt.Server
	Service *forward.Service
}

func NewForwardHandler(server *mqtt.Server, svc *forward.Service) *ForwardHandler {
	handler := &ForwardHandler{
		Server:  server,
		Service: svc,
	}
	server.AddHook(handler, nil)
	return handler
}

// ID returns the unique identifier of the hook.
func (h *ForwardHandler) ID() string {
	return "forward-hook"
}

// OnConnect is called when a client connects to the broker.
func (h *ForwardHandler) OnConnect(cl *mqtt.Client, pk packets.Packet) error {
	topic := pk.TopicName
	publisher := cl.ID
	h.Service.Connect(topic, publisher)
	return nil
}

// OnDisconnect is called when a client disconnects from the broker.
func (h *ForwardHandler) OnDisconnect(cl *mqtt.Client, err error, expire bool) {
	topic := cl.ID
	h.Service.Disconnect(topic)
}

// OnSubscribe is called when a client subscribes to a topic.
func (h *ForwardHandler) OnSubscribe(cl *mqtt.Client, pk packets.Packet) packets.Packet {
	id := cl.ID
	topic := pk.TopicName
	h.Service.Subscribe(topic, id)
	return pk
}

// OnUnsubscribe is called when a client unsubscribes from a topic.
func (h *ForwardHandler) OnUnsubscribe(cl *mqtt.Client, pk packets.Packet) packets.Packet {
	id := cl.ID
	topic := pk.TopicName
	h.Service.Unsubscribe(topic, id)
	return pk
}
