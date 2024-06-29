package kafka

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type (
	Deps struct {
		fx.In
		Config *viper.Viper
		Log    *log.Logger
	}

	Conn struct {
		log       *log.Logger
		consumer  *kafka.Consumer
		producer  *kafka.Producer
		consumers map[string][]consumerFunc
	}

	consumerFunc   = func(*kafka.Message)
	Message        = kafka.Message
	TopicPartition = kafka.TopicPartition
)

const (
	PartitionAny = kafka.PartitionAny
)

func (c *Conn) Run(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				msg, err := c.consumer.ReadMessage(100 * time.Millisecond)
				if err != nil {
					continue
				}

				consumers, ok := c.consumers[*msg.TopicPartition.Topic]
				if !ok {
					continue
				}

				var wg sync.WaitGroup
				for _, consumerFn := range consumers {
					consumerFn := consumerFn
					wg.Add(1)
					go func() {
						wg.Done()
						consumerFn(msg)
					}()

					wg.Wait()
				}
			}
		}
	}()

	go func() {
		for e := range c.producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					c.log.Printf("delivering message to partition: %v\n", ev.TopicPartition)
				}
			}
		}
	}()
}

func (c *Conn) Subscribe(topic string, consumer consumerFunc) {
	if c.consumers == nil {
		c.consumers = map[string][]consumerFunc{}
	}

	c.consumers[topic] = append(c.consumers[topic], consumer)
}

func (c *Conn) Publish(topic string, msg []byte) error {
	return c.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          msg,
	}, nil)
}

func (c *Conn) close() {
	c.producer.Flush(15 * 1000)
	c.producer.Close()
	c.consumer.Close()
}

func New(lc fx.Lifecycle, deps Deps) *Conn {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": deps.Config.GetString("kafka.bootstrap.servers"),
		"group.id":          deps.Config.GetString("app.name"),
		"auto.offset.reset": deps.Config.GetString("kafka.auto.offset.reset"),
	})

	if err != nil {
		panic(fmt.Sprintf("creating consumer: %v", err))
	}

	err = c.SubscribeTopics(deps.Config.GetStringSlice("kafka.topics"), nil)
	if err != nil {
		panic(fmt.Sprintf("subscribing consumer: %v", err))
	}

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": deps.Config.GetString("kafka.bootstrap.servers"),
		"acks":              deps.Config.GetString("kafka.acks"),
	})

	if err != nil {
		panic(fmt.Sprintf("creating producer: %v", err))
	}

	conn := &Conn{consumer: c, producer: p, log: deps.Log}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			conn.Run(ctx)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			conn.close()
			return nil
		},
	})

	return conn
}
