package hook

import (
	"log"

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

// Provides 声明该钩子提供哪些功能
func (h *ForwardHandler) Provides(b byte) bool {
	switch b {
	case mqtt.OnConnect, mqtt.OnDisconnect, mqtt.OnSubscribe, mqtt.OnUnsubscribe:
		return true
	}
	return false
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
	log.Println("OnConnect topic ", pk.TopicName, "client ", cl.ID)
	return nil
}

// OnDisconnect is called when a client disconnects from the broker.
func (h *ForwardHandler) OnDisconnect(cl *mqtt.Client, err error, expire bool) {
	topic := cl.ID
	h.Service.Disconnect(topic)
	log.Println("OnDisconnect client ", cl.ID)
}

// OnSubscribe is called when a client subscribes to a topic.
func (h *ForwardHandler) OnSubscribe(cl *mqtt.Client, pk packets.Packet) packets.Packet {
	id := cl.ID
	topic := pk.TopicName
	h.Service.Subscribe(topic, id)
	log.Println("OnSubscribe topic ", pk.TopicName, "client ", cl.ID)
	return pk
}

// OnUnsubscribe is called when a client unsubscribes from a topic.
func (h *ForwardHandler) OnUnsubscribe(cl *mqtt.Client, pk packets.Packet) packets.Packet {
	id := cl.ID
	topic := pk.TopicName
	h.Service.Unsubscribe(topic, id)
	log.Println("OnUnsubscribe topic ", pk.TopicName, "client ", cl.ID)
	return pk
}
