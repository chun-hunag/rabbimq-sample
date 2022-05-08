package rabbitMQService

import "github.com/streadway/amqp"

type PublishBuilder struct {
	publish amqp.Publishing
}

func (p *PublishBuilder) Init() *PublishBuilder {
	p.publish = amqp.Publishing{}
	return p
}

func (p PublishBuilder) Build() amqp.Publishing {
	return p.publish
}

func (p *PublishBuilder) SetDeliveryMode(deliveryMode uint8) *PublishBuilder {
	p.publish.DeliveryMode = deliveryMode
	return p
}

func (p *PublishBuilder) SetContentType(contentType string) *PublishBuilder {
	p.publish.ContentType = contentType
	return p
}

func (p *PublishBuilder) SetCorrelationId(correlationId string) *PublishBuilder {
	p.publish.CorrelationId = correlationId
	return p
}

func (p *PublishBuilder) SetReplyTo(replyTo string) *PublishBuilder {
	p.publish.ReplyTo = replyTo
	return p
}

func (p *PublishBuilder) SetBody(body []byte) *PublishBuilder {
	p.publish.Body = body
	return p
}
