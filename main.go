package main

type queueMethods interface {
	len() int
	first() *element
	last() *element
	equeue(value interface{})
	dequeue() interface{}
}

type element struct {
	previous *element
	next     *element
	value    interface{}
}

type queue struct {
	front *element
	back  *element
	size  int
}

func (receiver *queue) len() int {
	return receiver.size
}

func (receiver *queue) first() *element {
	return receiver.front
}

func (receiver queue) last() *element {
	return receiver.back
}

func (receiver *queue) equeue(value interface{}) {
	if receiver.size == 0 {
		receiver.size++
		receiver.front = &element{
			previous: nil,
			next:     nil,
			value:    value,
		}
		receiver.back = receiver.front
	} else {
		receiver.size++
		receiver.back.next = &element{
			previous: receiver.back,
			next:     nil,
			value:    value,
		}
		receiver.back = receiver.back.next
	}
}

func (receiver *queue) dequeue() interface{} {
	if receiver.size == 0 {
		return nil
	} else if receiver.size == 1 {
		currentValue := receiver.front.value
		receiver.front, receiver.back = nil, nil
		receiver.size--
		return currentValue
	} else {
		currentValue := receiver.front.value
		receiver.front = receiver.front.next
		receiver.front.previous = nil
		receiver.size--
		return currentValue
	}
}

func main() {}
