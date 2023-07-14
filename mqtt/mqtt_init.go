package mqtt

import (
	"AlliancesDocking/config"
	"AlliancesDocking/utils"
	"errors"
	"fmt"
	mymqtt "github.com/eclipse/paho.mqtt.golang"
	"sync"
	"time"
)

const (
	Topic_ping              = "Avtronsys/+/ping"
	topic_pong              = "Avtronsys/+/pong"
	topic_card_slots_change = "Avtronsys/+/card_slots_change"
	topic_paycard_local     = "Avtronsys/+/paycard_local"
	topic_paycard           = "Avtronsys/+/paycard"
	topic_get_device        = "Avtronsys/+/get_device"
	topic_device_push       = "Avtronsys/+/device_push"
	topic_device_recv       = "Avtronsys/+/device_recv"
	topic_fault_ctrl_panel  = "Avtronsys/+/fault_ctrl_panel"
	topic_card_logs         = "Avtronsys/+/card_logs"
	topic_update_req        = "Avtronsys/+/update_req"
	topic_download_ready    = "Avtronsys/+/download_ready"
	topic_download_update   = "Avtronsys/+/download_update"
	topic_card_change       = "Avtronsys/+/card_slots_change"
)

type Client struct {
	nativeClient  mymqtt.Client
	clientOptions *mymqtt.ClientOptions
	locker        *sync.Mutex
	// 消息收到之后处理函数
	observer func(c *Client, topic string, msg []byte)
}

var MqttClient *Client

func NewMqttClient(host string, clientId string, user string, pass string) *Client {
	clientOptions := mymqtt.NewClientOptions().
		AddBroker(host).
		SetUsername(user).
		SetPassword(pass).
		SetClientID(clientId).
		SetCleanSession(true).
		SetAutoReconnect(true).
		SetKeepAlive(120 * time.Second).
		SetPingTimeout(10 * time.Second).
		SetWriteTimeout(10 * time.Second).
		SetOnConnectHandler(func(client mymqtt.Client) {
			// 连接被建立后的回调函数
			fmt.Println("Mqtt is connected!", "clientId", clientId)
		}).
		SetConnectionLostHandler(func(client mymqtt.Client, err error) {
			// 连接被关闭后的回调函数
			fmt.Println("Mqtt is disconnected!", "clientId", clientId, "reason", err.Error())
		})

	nativeClient := mymqtt.NewClient(clientOptions)

	return &Client{
		nativeClient:  nativeClient,
		clientOptions: clientOptions,
		locker:        &sync.Mutex{},
	}
}

func MqttxServerInit() {
	clientID := fmt.Sprintf("avtronsysEDU_server_%s", utils.RandomString(5))
	MqttClient = NewMqttClient(config.GetConfig().Dev.Mqtt.Host, clientID,
		config.GetConfig().Dev.Mqtt.Username, config.GetConfig().Dev.Mqtt.Password)

	err := MqttClient.Connect() //连接mqtt,确保已经处于连接状态
	if err != nil {
		fmt.Println("mqtt connect error!!!")
		err = errors.New("mqtt connect error")
		return
	}
}

func (client *Client) Connect() error {
	return client.ensureConnected()
}

func (client *Client) ensureConnected() error {
	if !client.nativeClient.IsConnected() {
		client.locker.Lock()
		defer client.locker.Unlock()
		if !client.nativeClient.IsConnected() {
			if token := client.nativeClient.Connect(); token.Wait() && token.Error() != nil {
				return token.Error()
			}
		}
	}
	return nil
}

func (client *Client) Publish(topic string, qos byte, retained bool, data []byte) error {
	if err := client.ensureConnected(); err != nil {
		return err
	}

	token := client.nativeClient.Publish(topic, qos, retained, data)
	if err := token.Error(); err != nil {
		return err
	}

	// return false is the timeout occurred
	if !token.WaitTimeout(time.Second * 10) {
		return errors.New("mqtt publish wait timeout")
	}

	return nil
}

func (client *Client) Subscribe(observer func(c *Client, topic string, msg []byte), qos byte, topics ...string) error {
	if len(topics) == 0 {
		return errors.New("the topic is empty")
	}
	if observer == nil {
		return errors.New("the observer func is nil")
	}

	if client.observer != nil {
		return errors.New("an existing observer subscribed on this client, you must unsubscribe it before you subscribe a new observer")
	}
	client.observer = observer

	filters := make(map[string]byte)
	for _, topic := range topics {
		filters[topic] = qos
	}
	client.nativeClient.SubscribeMultiple(filters, client.messageHandler)

	return nil
}

func (client *Client) Unsubscribe(topics ...string) {
	client.observer = nil
	client.nativeClient.Unsubscribe(topics...)
}

func (client *Client) messageHandler(c mymqtt.Client, msg mymqtt.Message) {
	if client.observer == nil {
		//fmt.Println("not subscribe message observer")
		return
	}
	client.observer(client, msg.Topic(), msg.Payload())
}
