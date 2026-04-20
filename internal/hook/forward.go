package hook

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/bxcodec/go-clean-arch/forward"
	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
)

// ProtocolCallback 业务协议回调函数类型
type ProtocolCallback func(clientID string, protocol string, payload []byte) error

type ForwardHandler struct {
	mqtt.HookBase
	Server           *mqtt.Server
	Service          *forward.Service
	protocolHandlers map[string]ProtocolCallback
}

func NewForwardHandler(server *mqtt.Server, svc *forward.Service) *ForwardHandler {
	handler := &ForwardHandler{
		Server:           server,
		Service:          svc,
		protocolHandlers: make(map[string]ProtocolCallback),
	}
	server.AddHook(handler, nil)
	return handler
}

// RegisterProtocolHandler 注册协议处理器
func (h *ForwardHandler) RegisterProtocolHandler(protocol string, callback ProtocolCallback) {
	h.protocolHandlers[protocol] = callback
}

// Provides 声明该钩子提供哪些功能
func (h *ForwardHandler) Provides(b byte) bool {
	switch b {
	case mqtt.OnConnect, mqtt.OnDisconnect, mqtt.OnSubscribe, mqtt.OnUnsubscribe, mqtt.OnPublish:
		return true
	}
	return false
}

// ID returns the unique identifier of the hook.
func (h *ForwardHandler) ID() string {
	return "forward-hook"
}

// getClientName 获取客户端名称（优先使用 Username，否则使用 ID）
func getClientName(cl *mqtt.Client) string {
	if len(cl.Properties.Username) > 0 {
		return string(cl.Properties.Username)
	}
	return cl.ID
}

// buildPubTopic 构建 pub topic，格式: 名字/ID_pub
func buildPubTopic(cl *mqtt.Client) string {
	name := getClientName(cl)
	if name == cl.ID {
		// 如果没有单独的名字，使用 ID_pub 格式
		return cl.ID + "_pub"
	}
	return name + "_pub"
}

// OnConnect 当客户端连接时，自动订阅其 pub topic
func (h *ForwardHandler) OnConnect(cl *mqtt.Client, pk packets.Packet) error {
	clientID := cl.ID
	clientName := getClientName(cl)
	pubTopic := buildPubTopic(cl)

	// 保存 pub topic 信息
	h.Service.Connect(pubTopic, clientID)

	// 订阅该客户端的 pub topic（模拟订阅）
	h.Service.Subscribe(pubTopic, clientID)

	log.Printf("OnConnect: client=%s, name=%s, pub_topic=%s", clientID, clientName, pubTopic)
	return nil
}

// OnDisconnect 当客户端断开时，清理其 pub topic
func (h *ForwardHandler) OnDisconnect(cl *mqtt.Client, err error, expire bool) {
	clientID := cl.ID
	pubTopic := buildPubTopic(cl)

	// 取消订阅
	h.Service.Unsubscribe(pubTopic, clientID)

	// 删除 pub topic
	h.Service.Disconnect(pubTopic)

	log.Printf("OnDisconnect: client=%s, pub_topic=%s", clientID, pubTopic)
}

// OnSubscribe 当客户端订阅主题时
func (h *ForwardHandler) OnSubscribe(cl *mqtt.Client, pk packets.Packet) packets.Packet {
	// 处理多个订阅过滤器
	for _, filter := range pk.Filters {
		topic := filter.Filter
		clientID := cl.ID
		h.Service.Subscribe(topic, clientID)
		log.Printf("OnSubscribe: topic=%s, client=%s", topic, clientID)
	}
	return pk
}

// OnUnsubscribe 当客户端取消订阅时
func (h *ForwardHandler) OnUnsubscribe(cl *mqtt.Client, pk packets.Packet) packets.Packet {
	// 处理多个取消订阅过滤器
	for _, filter := range pk.Filters {
		topic := filter.Filter
		clientID := cl.ID
		h.Service.Unsubscribe(topic, clientID)
		log.Printf("OnUnsubscribe: topic=%s, client=%s", topic, clientID)
	}
	return pk
}

// OnPublish 当收到消息时，解析 protocol 并进行回调
func (h *ForwardHandler) OnPublish(cl *mqtt.Client, pk packets.Packet) (packets.Packet, error) {
	clientID := cl.ID
	clientName := getClientName(cl)
	topic := pk.TopicName
	payload := pk.Payload

	log.Printf("OnPublish: topic=%s, client=%s, payload=%s", topic, clientID, string(payload))

	// 解析 JSON 获取 protocol 字段
	protocol, err := extractProtocol(payload)
	if err != nil {
		log.Printf("OnPublish: failed to extract protocol from client=%s: %v", clientID, err)
		return pk, nil // 不阻断消息，继续处理
	}

	log.Printf("OnPublish: client=%s, name=%s, protocol=%s", clientID, clientName, protocol)

	// 调用注册的回调
	if callback, ok := h.protocolHandlers[protocol]; ok {
		if err := callback(clientID, protocol, payload); err != nil {
			log.Printf("OnPublish: callback error for protocol=%s: %v", protocol, err)
		}
	} else {
		log.Printf("OnPublish: no handler registered for protocol=%s", protocol)
	}

	return pk, nil
}

// extractProtocol 从 JSON payload 中提取 protocol 字段
func extractProtocol(payload []byte) (string, error) {
	var data map[string]interface{}
	if err := json.Unmarshal(payload, &data); err != nil {
		return "", err
	}

	protocol, ok := data["protocol"].(string)
	if !ok {
		// 尝试从 Protocol（大写）获取
		protocol, ok = data["Protocol"].(string)
		if !ok {
			return "", nil
		}
	}

	return strings.TrimSpace(protocol), nil
}
