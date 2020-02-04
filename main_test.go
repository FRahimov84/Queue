package main

import (
	"testing"
)

func Test_queue_len(t *testing.T) {
	q := queue{}
	if q.len() != 0 {
		want := 0
		got := q.len()
		t.Errorf("method len() is not correct, for empty queue got: %v want: %v", got, want)
	}
}

func Test_queue_first(t *testing.T) {
	q := queue{}
	if got := q.first(); got != nil {
		t.Errorf("method first() is not correct, for empty queue got: %v want: %v", got, nil)
	}
	want := 1
	q.equeue(want)
	if got := q.first().value; got != want {
		t.Errorf("method first() is not correct, for queue with one element got: %v want: %v", got, want)
	}
}

func Test_queue_last(t *testing.T) {
	q := queue{}
	if got := q.last(); got != nil {
		t.Errorf("method last() is not correct, for empty queue got: %v want: %v", got, nil)
	}
	q.equeue(1)
	q.equeue(2)
	q.equeue(3)
	if got, want := q.last().value, 3; got != want {
		t.Errorf("method last() is not correct, for queue with three element got: %v want: %v", got, want)
	}
	if got, want:= q.len(), 3; got != want {
		t.Errorf("method len() is not correct, for queue with three element got: %v want: %v", got, want)
	}

	if got, want := q.first().value, 1; got != want {
		t.Errorf("method first() is not correct, for queue with three element got: %v want: %v", got, want)
	}
}

func Test_queue_equeue(t *testing.T) {
	q := queue{}
	q.equeue("first")
	if got, want := q.last().value, "first"; got != want {
		t.Errorf("method last() is not correct, for queue with one element got: %v want: %v", got, want)
	}
	if got, want:= q.len(), 1; got != want {
		t.Errorf("method len() is not correct, for queue with one element got: %v want: %v", got, want)
	}
	if got, want := q.first().value, "first"; got != want {
		t.Errorf("method first() is not correct, for queue with one element got: %v want: %v", got, want)
	}
	q.equeue("second")
	q.equeue("third")
	q.equeue("fourth")

	if got, want := q.last().value, "fourth"; got != want {
		t.Errorf("method last() is not correct, for queue with elements got: %v want: %v", got, want)
	}
	if got, want:= q.len(), 4; got != want {
		t.Errorf("method len() is not correct, for queue with elements got: %v want: %v", got, want)
	}

	if got, want := q.first().value, "first"; got != want {
		t.Errorf("method first() is not correct, for queue with elements got: %v want: %v", got, want)
	}
}

func Test_queue_dequeue(t *testing.T) {
	q := queue{}
	if got:=q.dequeue(); got!=nil{
		t.Errorf("method dequeue() is not correct, for empty queue got: %v want: %v", got, nil)
	}
	q.equeue("first")
	q.equeue("second")
	q.equeue("third")
	q.equeue("fourth")
	if got, want:= q.len(), 4; got != want {
		t.Errorf("method len() is not correct, for queue with elements got: %v want: %v", got, want)
	}
	if got, want:=q.dequeue(),"first"; got!=want{
		t.Errorf("method dequeue() is not correct, for queue with elements got: %v want: %v", got, want)
	}
	if got, want:= q.len(), 3; got != want {
		t.Errorf("method len() is not correct, for queue with elements got: %v want: %v", got, want)
	}
	if got, want := q.first().value, "second"; got != want {
		t.Errorf("method first() is not correct, for queue with elements got: %v want: %v", got, want)
	}
	if got, want := q.last().value, "fourth"; got != want {
		t.Errorf("method last() is not correct, for queue with elements got: %v want: %v", got, want)
	}
	q.dequeue()
	q.dequeue()
	q.dequeue()
	if got, want:= q.len(), 0; got != want {
		t.Errorf("method len() is not correct, for empty queue got: %v want: %v", got, want)
	}
	if got:= q.first(); got != nil {
		t.Errorf("method first() is not correct, for queue with elements got: %v want: %v", got, nil)
	}
	if got := q.last(); got != nil {
		t.Errorf("method last() is not correct, for queue with elements got: %v want: %v", got, nil)
	}
}