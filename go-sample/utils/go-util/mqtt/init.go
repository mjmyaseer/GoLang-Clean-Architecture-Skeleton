package mqtt

import (
	PahoMqtt "github.com/eclipse/paho.mqtt.golang"
	"go-sample/utils/go-util/config"
	"go-sample/utils/go-util/log"
	"os"
	"os/signal"
	"time"
)

var (
	opts   *PahoMqtt.ClientOptions
	client PahoMqtt.Client
)

type conf struct {
	Brokers              []string `yaml:"brokers" json:"brokers"`
	ClientID             string   `yaml:"client_id" json:"client_id"`
	PingTimeout          int      `yaml:"ping_timeout" json:"ping_timeout"`
	MaxReconnectInterval int      `yaml:"max_reconnect_interval" json:"max_reconnect_interval"`
	ConnectTimeout       int      `yaml:"connect_timeout" json:"connect_timeout"`
}

var mqttConf conf

func Init(clientId string, onConnect PahoMqtt.OnConnectHandler) {
	conf := parseConfig()
	opts = PahoMqtt.NewClientOptions()
	opts.CleanSession = false
	opts.AutoReconnect = true
	opts.PingTimeout = time.Duration(conf.PingTimeout) * time.Second
	opts.MaxReconnectInterval = time.Duration(conf.MaxReconnectInterval) * time.Second
	opts.ConnectTimeout = time.Duration(conf.ConnectTimeout) * time.Second
	opts.OnConnectionLost = func(client PahoMqtt.Client, e error) {
		log.Error(`mqtt client disconnected `, e)
	}
	opts.OnConnect = onConnect
	for _, addr := range conf.Brokers {
		opts.AddBroker(`tcp://` + addr)
	}

	opts.ClientID = conf.ClientID
	if clientId != `` {
		opts.ClientID = clientId
	}

	client = PahoMqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Error(`Cannot connect to the broker : `, token.Error())
		return
	}

	log.Info(`MQTT Connection establish for client `, opts.ClientID)

	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt)

		select {
		case sig := <-signals:
			client.Disconnect(200)
			log.Info(`Mqtt connection aborted : `, sig)
			break
		}
	}()
}

func parseConfig() (conf conf) {
	config.DefaultConfigurator.Load(`config/mqtt`, &mqttConf, func(config interface{}) {

	})

	return mqttConf
}
