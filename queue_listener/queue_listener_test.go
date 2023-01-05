package queue_listener_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/william22913/common/queue"
	"github.com/william22913/common/queue_listener"
)

func TestSomething(t *testing.T) {

	queue := queue.NewQueue()
	queueListener := queue_listener.NewQueueListener(queue, 100)
	queueListener.ListenQueue(doFunction)
	defer queueListener.StopListen()

	for i := 0; i < 10; i++ {
		queue.Push([]byte{100})
		time.Sleep(1 * time.Second)
	}

}

func doFunction(x int, input interface{}) {
	fmt.Println("thread [", x, "] processing message", "->"+string(input.([]byte))+"<-", "at", time.Now())
	time.Sleep(2 * time.Second)
}
