package messaging

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// PedidoQueue representa el pedido que se enviará a la cola
type PedidoQueue struct {
	ID            uint    `json:"id"`
	UsuarioID     uint    `json:"usuario_id"`
	RestauranteID uint    `json:"restaurante_id"`
	Producto1     string  `json:"producto1"`
	Producto2     string  `json:"producto2"`
	Producto3     string  `json:"producto3"`
	Total         float64 `json:"total"`
	Estado        string  `json:"estado"`
	PagoID        uint    `json:"pago_id"`
}

// PedidoPublisher es el publicador de mensajes a RabbitMQ
type PedidoPublisher struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
}

// NewPedidoPublisher inicializa el publicador
func NewPedidoPublisher() (*PedidoPublisher, error) {
	conn, err := amqp.Dial("amqp://dvelazquez:laconia@75.101.219.208:5672/")
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		"pedidos", // Nombre de la cola
		true,      // Durable
		false,     // Auto-delete
		false,     // Exclusive
		false,     // No-wait
		nil,       // Arguments
	)
	if err != nil {
		return nil, err
	}

	return &PedidoPublisher{
		conn:    conn,
		channel: ch,
		queue:   q,
	}, nil
}

// Publish envía un pedido a la cola de RabbitMQ
func (p *PedidoPublisher) Publish(pedido PedidoQueue) error {
	body, err := json.Marshal(pedido)
	if err != nil {
		return err
	}

	err = p.channel.Publish(
		"",           
		p.queue.Name, 
		false,        
		false,        
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return err
	}

	log.Printf("Pedido enviado a la cola: %+v\n", pedido)
	return nil
}

// Close cierra la conexión y el canal de RabbitMQ
func (p *PedidoPublisher) Close() {
	p.channel.Close()
	p.conn.Close()
}
