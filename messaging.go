package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Producer interface {
	connect()
	setup()
	publish()
}

type Subscriber interface {
	connect()
	setup()
	subscribe()
}

type ExchangeSettings struct {
	exchangeName       string
	exchangeType       string
	exchangeDurable    bool
	exchangeAutoDelete bool
	exchangeInternal   bool
	exchangeNoWait     bool
	exchangeArguments  amqp.Table
}

type RabbitMQService struct {
	c  amqp.Connection
	ch amqp.Channel
	ac AppConfig
}

func (rmq *RabbitMQService) connect() {
	conn, err := amqp.Dial(rmq.ac.rabbitmqConnectionString)

	if err != nil {
		log.Fatalln(err)
	}

	rmq.c = *conn
}

func (rmq *RabbitMQService) createChannel() {
	channel, err := rmq.c.Channel()

	if err != nil {
		log.Fatalln(err)
	}

	rmq.ch = *channel
}

func (rmq *RabbitMQService) createTopic(es ExchangeSettings) {
	err := rmq.ch.ExchangeDeclare(
		es.exchangeName,
		es.exchangeType,
		es.exchangeDurable,
		es.exchangeAutoDelete,
		es.exchangeInternal,
		es.exchangeNoWait,
		es.exchangeArguments,
	)

	if err != nil {
		log.Fatalln(err)
	}
}

type RMQMessageConfig struct {
	exchangeName       string
	routingKey         string
	messageMandatory   bool
	messageImmediate   bool
	messageContentType string
}

func (rmq RabbitMQService) publishMessage(rmqMessage RMQMessageConfig, body []byte) {
	publishingConfig := amqp.Publishing{
		ContentType: rmqMessage.messageContentType,
		Body:        body,
	}

	err := rmq.ch.Publish(
		rmqMessage.exchangeName,
		rmqMessage.routingKey,
		rmqMessage.messageMandatory,
		rmqMessage.messageImmediate,
		publishingConfig,
	)

	if err != nil {
		log.Fatalln(err)
	}
}
