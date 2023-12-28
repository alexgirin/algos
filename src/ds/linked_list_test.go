package ds

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedListGet(t *testing.T) {
	var val int
	var err error
	list := NewLinkedList[int]()

	val, err = list.GetFirst()
	assert.NotNil(t, err)
	assert.Equal(t, 0, val)

	val, err = list.GetLast()
	assert.NotNil(t, err)
	assert.Equal(t, 0, val)

	list.Push(10)
	val, err = list.GetFirst()
	assert.Nil(t, err)
	assert.Equal(t, 10, val)

	list.Push(20)
	list.Push(30)
	val, err = list.GetLast()
	assert.Nil(t, err)
	assert.Equal(t, 30, val)
}

func TestLinkedListPush(t *testing.T) {
	list := NewLinkedList[int]()

	list.Push(10)
	assert.Equal(t, 1, list.Size())

	list.Push(20)
	assert.Equal(t, 2, list.Size())
}

func TestLinkedListPop(t *testing.T) {
	var val int
	var err error
	list := NewLinkedList[int]()

	list.Push(10)
	list.Push(20)
	assert.Equal(t, 2, list.Size())

	val, err = list.Pop()
	assert.Nil(t, err)
	assert.Equal(t, 20, val)
	assert.Equal(t, 1, list.Size())

	val, err = list.Pop()
	assert.Nil(t, err)
	assert.Equal(t, 10, val)
	assert.Equal(t, 0, list.Size())

	val, err = list.Pop()
	assert.NotNil(t, err)
	assert.Equal(t, 0, val)
}

func TestLinkedListEnqueue(t *testing.T) {
	list := NewLinkedList[int]()

	list.Enqueue(10)
	assert.Equal(t, 1, list.Size())

	list.Enqueue(20)
	assert.Equal(t, 2, list.Size())

	val, err := list.GetFirst()
	assert.Nil(t, err)
	assert.Equal(t, 20, val)
}

func TestLinkedListDequeue(t *testing.T) {
	var val int
	var err error
	list := NewLinkedList[int]()

	list.Push(10)
	list.Push(20)

	val, err = list.Dequeue()
	assert.Nil(t, err)
	assert.Equal(t, 10, val)
	assert.Equal(t, 1, list.Size())

	val, err = list.Dequeue()
	assert.Nil(t, err)
	assert.Equal(t, 20, val)
	assert.Equal(t, 0, list.Size())

	val, err = list.Dequeue()
	assert.NotNil(t, err)
	assert.Equal(t, 0, val)
}

func TestLinkedListContains(t *testing.T) {
	list := NewLinkedList[int]()

	assert.False(t, list.Contains(0))

	list.Push(10)
	list.Push(20)

	assert.False(t, list.Contains(0))
	assert.True(t, list.Contains(10))
	assert.True(t, list.Contains(20))
	assert.False(t, list.Contains(30))
}

func TestLinkedListToSlice(t *testing.T) {
	list := NewLinkedList[int]()
	assert.True(t, reflect.DeepEqual(list.ToSlice(), []int{}))

	list.Push(10)
	list.Push(20)
	assert.True(t, reflect.DeepEqual(list.ToSlice(), []int{10, 20}))

	list.Enqueue(30)
	assert.True(t, reflect.DeepEqual(list.ToSlice(), []int{30, 10, 20}))
}

func TestLinkedListInsert(t *testing.T) {
	var err error
	list := NewLinkedList[int]()

	err = list.Insert(1, 10)
	assert.NotNil(t, err)

	err = list.Insert(0, 10)
	assert.Nil(t, err)
	assert.Equal(t, 1, list.Size())

	list.Push(20)
	list.Push(30)

	err = list.Insert(1, 40)
	assert.Nil(t, err)
	assert.Equal(t, 4, list.Size())

	err = list.Insert(4, 50)
	assert.Nil(t, err)
	assert.Equal(t, 5, list.Size())
	assert.True(t, reflect.DeepEqual(list.ToSlice(), []int{10, 40, 20, 30, 50}))

	err = list.Insert(6, 60)
	assert.NotNil(t, err)
}
