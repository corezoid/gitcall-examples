package main

import (
	"context"
	"log"

	"encoding/json"

	"github.com/corezoid/gitcall-go-runner/gitcall"
	"github.com/streadway/amqp"
)

var amqpDSN = "amqp://guest:guest@localhost:5672/%2f"
var _ch *amqp.Channel

func usercode(_ context.Context, data map[string]interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = ch().Publish("", "gitcall-example", false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        b,
	})
	if err != nil {
		return err
	}

	return nil
}

func main() {
	gitcall.Handle(usercode)
}

func ch() *amqp.Channel {
	if _ch == nil {
		conn, err := amqp.Dial(amqpDSN)
		if err != nil {
			log.Fatal(err)
		}

		_ch, err = conn.Channel()
		if err != nil {
			log.Fatal(err)
		}

		_, err = _ch.QueueDeclare("gitcall-example", true, false, false, false, nil)
		if err != nil {
			log.Fatal(err)
		}
	}

	return _ch
}
