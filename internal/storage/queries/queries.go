package queries

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/imirjar/rb-michman/internal/models"
	"github.com/streadway/amqp"
)

type MQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func New() *MQ {
	return &MQ{}
}

func (m *MQ) Connect(ctx context.Context, connection string) error {
	conn, err := amqp.Dial(connection)
	if err != nil {
		return fmt.Errorf("failed to connect to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return fmt.Errorf("failed to create channel: %v", err)
	}

	m.conn = conn
	m.channel = ch

	// Объявляем очередь и exchange
	_, err = ch.QueueDeclare(
		"data_queue",
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %v", err)
	}

	return nil
}

func (m *MQ) Disconnect() error {
	if m.channel != nil {
		if err := m.channel.Close(); err != nil {
			log.Printf("Failed to close channel: %v", err)
		}
	}
	if m.conn != nil {
		if err := m.conn.Close(); err != nil {
			return fmt.Errorf("failed to close connection: %v", err)
		}
	}
	return nil
}

func (m *MQ) ExecuteReport(ctx context.Context, repId string) (models.Data, error) {
	if m.channel == nil {
		return models.Data{}, fmt.Errorf("RabbitMQ channel not initialized")
	}

	// Создаем временную очередь для ответа
	replyQueue, err := m.channel.QueueDeclare(
		"",    // случайное имя
		false, // durable
		true,  // autoDelete
		true,  // exclusive
		false, // noWait
		nil,   // args
	)
	if err != nil {
		return models.Data{}, fmt.Errorf("failed to declare reply queue: %v", err)
	}

	// Подписываемся на очередь ответов
	msgs, err := m.channel.Consume(
		replyQueue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return models.Data{}, fmt.Errorf("failed to consume reply queue: %v", err)
	}

	correlationId := generateCorrelationId()

	// Публикуем запрос
	err = m.channel.Publish(
		"",
		"data_queue",
		false,
		false,
		amqp.Publishing{
			ContentType:   "text/plain",
			Body:          []byte(repId),
			ReplyTo:       replyQueue.Name,
			CorrelationId: correlationId,
		},
	)
	if err != nil {
		return models.Data{}, fmt.Errorf("failed to publish message: %v", err)
	}

	// Ждем ответа с таймаутом
	select {
	case msg := <-msgs:
		if msg.CorrelationId != correlationId {
			return models.Data{}, fmt.Errorf("correlation ID mismatch")
		}

		var data models.Data
		if err := json.Unmarshal(msg.Body, &data); err != nil {
			return models.Data{}, fmt.Errorf("failed to unmarshal response: %v", err)
		}
		return data, nil

	case <-time.After(30 * time.Second):
		return models.Data{}, fmt.Errorf("timeout waiting for response")

	case <-ctx.Done():
		return models.Data{}, ctx.Err()
	}
}

func generateCorrelationId() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
