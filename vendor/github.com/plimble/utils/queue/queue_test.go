package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestLengthZero(t *testing.T) {
	q := NewQueue()
	assert.Equal(t, 0, q.Len())
}

//--Test Push

func TestSimplePoll(t *testing.T) {
	q := NewQueue()
	expected := 20
	q.Push(expected)
	assert.Equal(t, 1, q.Len())
	val := q.Peek()

	assert.Equal(t, expected, val)
	assert.Equal(t, 1, q.Len())
}

func TestPushNils(t *testing.T) {
	q := NewQueue()
	for i := 0; i < 15; i++ {
		q.Push(nil)
		assert.Equal(t, i+1, q.Len())
	}
	assert.Equal(t, 15, q.Len())
	for i := 15; i > 0; i-- {
		assert.Equal(t, i, q.Len())
		val := q.Poll()
		assert.Equal(t, nil, val)
		assert.Equal(t, i-1, q.Len())
	}
}

func TestMixedPushes(t *testing.T) {
	q := NewQueue()
	assert.Equal(t, 0, q.Len())

	//Pushes
	q.Push(nil)
	assert.Equal(t, 1, q.Len())

	q.Push(10)
	assert.Equal(t, 2, q.Len())

	q.Push("foo")
	assert.Equal(t, 3, q.Len())

	q.Push([]int{1, 2, 3})
	assert.Equal(t, 4, q.Len())

	q.Push(nil)
	assert.Equal(t, 5, q.Len())

	q.Push("bar")
	assert.Equal(t, 6, q.Len())

	q.Push(nil)
	assert.Equal(t, 7, q.Len())

	q.Push(nil)
	assert.Equal(t, 8, q.Len())

	//Polls
	var val interface{}
	val = q.Poll()
	assert.Equal(t, 7, q.Len())
	assert.Nil(t, val)

	val = q.Poll()
	assert.Equal(t, 6, q.Len())
	assert.Equal(t, 10, val)

	val = q.Poll()
	assert.Equal(t, 5, q.Len())
	assert.Equal(t, "foo", val)

	val = q.Poll()
	assert.Equal(t, 4, q.Len())
	assert.Equal(t, []int{1, 2, 3}, val)

	val = q.Poll()
	assert.Equal(t, 3, q.Len())
	assert.Nil(t, val)
	assert.Equal(t, 3, q.Len())
	assert.Nil(t, val)

	val = q.Poll()
	assert.Equal(t, 2, q.Len())
	assert.Equal(t, "bar", val)

	val = q.Poll()
	assert.Equal(t, 1, q.Len())
	assert.Nil(t, val)

	val = q.Poll()
	assert.Equal(t, 0, q.Len())
	assert.Nil(t, val)
}

//--Test Poll

func TestPullNothing(t *testing.T) {
	q := NewQueue()
	val := q.Poll()
	assert.Nil(t, val)
	assert.Equal(t, 0, q.Len())
}

func TestMultiplePolls(t *testing.T) {
	q := NewQueue()
	for i := 0; i < 10; i++ {
		q.Push(i)
		assert.Equal(t, 0, q.Peek())
		assert.Equal(t, i+1, q.Len())
	}
	i := 0
	j := 10
	for i < 10 {
		length := j
		assert.Equal(t, length, q.Len())
		val := q.Poll()
		assert.Equal(t, i, val)
		assert.Equal(t, length-1, q.Len())
		i++
		j--
	}
	assert.Equal(t, 0, q.Len())

	v := q.Poll()
	assert.Nil(t, v)
	assert.Equal(t, 0, q.Len())

	v = q.Poll()
	assert.Nil(t, v)
	assert.Equal(t, 0, q.Len())

	for i := 0; i < 20; i++ {
		v = q.Poll()
	}
	assert.Nil(t, v)
	assert.Equal(t, 0, q.Len())
}

//--Peeks tests

func TestReadStuf(t *testing.T) {
	q := NewQueue()
	expected := 20
	q.Push(expected)
	assert.Equal(t, 1, q.Len())
	val := q.Peek()
	assert.Equal(t, expected, val)
	assert.Equal(t, 1, q.Len())
}

//--Concurrent tests

// func TestConcurrent(t *testing.T) {
// 	q := NewQueue()
// 	sleepTime := 100
// 	numberGoRoutines := 50
// 	numberOfPushes := 10000
// 	ch := make(chan int, numberGoRoutines)
// 	for i := 0; i < numberGoRoutines; i++ {
// 		go inceremtQueue(q, ch, numberOfPushes)
// 	}
// 	Wait(sleepTime)
// 	for {
// 		if len(ch) == numberGoRoutines {
// 			assert.Equal(t, numberGoRoutines, len(ch))
// 			assert.Equal(t, numberGoRoutines*numberOfPushes, q.Len())
// 			return
// 		}
// 	}
// }

func inceremtQueue(q *Queue, ch chan<- int, numberOfPushes int) {
	for j := 0; j < numberOfPushes; j++ {
		q.Push(j)
	}
	ch <- 1
}

func Wait(duration int) {
	time.Sleep(time.Duration(duration) * time.Millisecond)
}
