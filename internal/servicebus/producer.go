package servicebus

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"github.com/google/uuid"
)

type Producer struct {
	client     *azservicebus.Client
	queueName  string
	replyQueue string
}

func NewProducer(conn, queue, replyQueue string) *Producer {
	client, err := azservicebus.NewClientFromConnectionString(conn, nil)
	if err != nil {
		panic(err)
	}

	return &Producer{
		client:     client,
		queueName:  queue,
		replyQueue: replyQueue,
	}
}

func (p *Producer) Send(message string) (string, error) {
	ctx := context.Background()

	sender, err := p.client.NewSender(p.queueName, nil)
	if err != nil {
		return "", err
	}
	defer sender.Close(ctx)

	correlationId := uuid.NewString()

	msg := &azservicebus.Message{
		Body:          []byte(message),
		CorrelationID: &correlationId,
		ReplyTo:       &p.replyQueue,
	}

	err = sender.SendMessage(ctx, msg, nil)
	if err != nil {
		return "", err
	}

	receiver, err := p.client.NewReceiverForQueue(p.replyQueue, nil)
	if err != nil {
		return "", err
	}
	defer receiver.Close(ctx)

	for {
		messages, err := receiver.ReceiveMessages(ctx, 1, nil)
		if err != nil {
			return "", err
		}

		for _, message := range messages {
			if message.CorrelationID != nil && *message.CorrelationID == correlationId {
				err = receiver.CompleteMessage(ctx, message, nil)
				if err != nil {
					return "", err
				}
				return string(message.Body), nil
			}
		}
	}

}
