package servicebus

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
)

type Producer struct {
	client    *azservicebus.Client
	queueName string
}

func NewProducer(conn string, queue string) *Producer {
	client, err := azservicebus.NewClientFromConnectionString(conn, nil)
	if err != nil {
		panic(err)
	}

	return &Producer{
		client:    client,
		queueName: queue,
	}
}

func (p *Producer) Send(message string) error {
	ctx := context.Background()

	sender, err := p.client.NewSender(p.queueName, nil)
	if err != nil {
		return err
	}

	defer sender.Close(ctx)

	msg := &azservicebus.Message{
		Body: []byte(message),
	}

	return sender.SendMessage(ctx, msg, nil)

}
