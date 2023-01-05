package queue_listener

import (
	"time"

	"github.com/william22913/common/queue"
)

type queueListener struct {
	queue              queue.Queue
	number_of_listener int
	state              []chan struct{}
}

func (q *queueListener) ListenQueue(f DoFunction) {
	for i := 0; i < q.number_of_listener; i++ {
		q.state = append(q.state, q.startListen(i, f))
	}
}

func (q *queueListener) startListen(
	x int,
	f DoFunction,
) chan struct{} {
	state := make(chan struct{})

	go func() {
		for {
			select {
			case <-time.After(1 * time.Second):
				//TODO: Process your job here
				msg := q.queue.Pop()
				if msg == nil {
					continue
				}

				f(x, msg)
			case <-state:
				break
			}
		}
	}()

	return state

}

func (q *queueListener) StopListen() {
	for i := 0; i < len(q.state); i++ {
		q.state[i] <- struct{}{}
	}
}
