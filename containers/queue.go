package containers

import (
	"container/list"
	"fmt"
)

type customQueue struct {
	queue *list.List
}

func (c *customQueue) Enqueue(value string) {
	c.queue.PushBack(value)
}

func (c *customQueue) Dequeue() error {
	if c.queue.Len() > 0 {
		ele := c.queue.Front()
		c.queue.Remove(ele)
		return nil
	}

	return fmt.Errorf("pop error: queue is empty")
}

func (c *customQueue) Front() (string, error) {
	if c.queue.Len() > 0 {
		if val, ok := c.queue.Front().Value.(string); ok {
			return val, nil
		}

		return "", fmt.Errorf("peep error: queue datatype is incorrect")
	}
	return "", fmt.Errorf("peep error: queue is empty")
}

func (c *customQueue) Size() int {
	return c.queue.Len()
}

func (c *customQueue) Empty() bool {
	return c.queue.Len() == 0
}

func TestQueue() {
	customQueue := &customQueue{
		queue: list.New(),
	}

	customQueue.Enqueue("A")
	customQueue.Enqueue("B")
	fmt.Printf("Size: %d\n", customQueue.Size())

	for customQueue.Size() > 0 {
		frontVal, _ := customQueue.Front()
		fmt.Printf("Front: %s\n", frontVal)
		customQueue.Dequeue()
	}

	fmt.Printf("Size: %d\n", customQueue.Size())
}
