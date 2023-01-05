package queue_listener

import "github.com/william22913/common/queue"

type DoFunction func(int, interface{})

func NewQueueListener(
	queue queue.Queue,
	number_of_listener int,
) QueueListener {
	return &queueListener{
		queue:              queue,
		number_of_listener: number_of_listener,
	}
}

type QueueListener interface {
	ListenQueue(DoFunction)

	StopListen()
}
