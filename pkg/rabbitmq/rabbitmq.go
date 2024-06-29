package rabbitmq

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type (
	Deps struct {
		fx.In
		Config *viper.Viper
		Logger *log.Logger
	}

	Conn struct {
		conn *amqp.Connection
	}

	consumerFunc = func(amqp.Delivery)
)

func New(deps Deps) *Conn {
	conn, err := amqp.Dial(deps.Config.GetString("rabbitmq.url"))
	if err != nil {
		deps.Logger.Fatalf("connecting rabbitmq: %v", err)
	}

	return &Conn{conn: conn}
}

func (c *Conn) Channel() (*amqp.Channel, error) {
	return c.conn.Channel()
}

func (c *Conn) Subscribe(queue string, consumer consumerFunc) error {
	rxCh, err := c.conn.Channel()
	if err != nil {
		return fmt.Errorf("creating amqp rx channel: %v", err)
	}

	msgs, err := rxCh.Consume(queue, "", false, false, false, false, nil)
	if err != nil {
		return fmt.Errorf("consuming amqp messages: %v", err)
	}

	go func() {
		for d := range msgs {
			consumer(d)
		}
	}()

	return nil
}
