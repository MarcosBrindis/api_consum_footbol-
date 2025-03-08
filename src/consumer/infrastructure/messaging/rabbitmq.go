package messaging

import (
	"context"
	"errors"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
}

func NewRabbitMQ(url, queueName string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		queueName, // nombre de la cola
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return nil, err
	}

	return &RabbitMQ{
		conn:    conn,
		channel: ch,
		queue:   q,
	}, nil
}

// GetMessage obtiene un mensaje de la cola y lo elimina (haciendo ACK)
func (r *RabbitMQ) GetMessage(ctx context.Context) (string, error) {
	msg, ok, err := r.channel.Get(r.queue.Name, false) // false para no auto-ack
	if err != nil {
		return "", err
	}
	if !ok {
		return "", errors.New("no hay mensajes en la cola")
	}
	// Confirmamos (ACK) la recepción para eliminar el mensaje de la cola
	if err := r.channel.Ack(msg.DeliveryTag, false); err != nil {
		return "", err
	}
	log.Printf("Mensaje consumido: %s", string(msg.Body))
	return string(msg.Body), nil
}

func (r *RabbitMQ) Close() {
	r.channel.Close()
	r.conn.Close()
}

// Función de inicialización, ajusta la URL y el nombre de la cola según corresponda
func InitRabbitMQ() (*RabbitMQ, error) {
	return NewRabbitMQ("amqp://MarcosDaniel:123456789a@54.146.109.24:5672/", "sport_events")
}
