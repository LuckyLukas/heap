package binaryheap

import "testing"
import "github.com/stretchr/testify/assert"


type IntNode struct {
	content int64
}

func (i IntNode) Value() int64 {
	return i.content
}

func TestNew(t *testing.T) {
	h := New(0)
	assert.NotNil(t, h)
}

func TestHeap_Add(t *testing.T) {
	h := New(0)
	node := IntNode{2}
	h.Add(node)
	actual, _ := h.Top()
	assert.Equal(t, node, actual)
}

func TestHeap_Add2(t *testing.T) {
	h := New(0)
	bad := IntNode{4}
	better := IntNode{2}
	worse := IntNode{5}
	h.Add(bad)
	h.Add(better)
	h.Add(worse)
	actual, _ := h.Top()
	assert.Equal(t, better.Value(), actual.Value())
}

func TestHeap_Remove(t *testing.T) {
	h := New(0)
	bad := &IntNode{4}
	better := &IntNode{2}
	worse := &IntNode{5}
	h.Add(bad)
	h.Add(better)
	h.Add(worse)

	top, err  :=  h.Top()
	assert.NoError(t, err)
	actual, err := h.RemoveTop()
	assert.NoError(t, err)
	assert.Equal(t, actual.Value(), top.Value())
	assert.Equal(t, better.Value(), actual.Value())

	top, err  =  h.Top()
	assert.NoError(t, err)
	actual, err = h.RemoveTop()
	assert.NoError(t, err)
	assert.Equal(t, actual.Value(), top.Value())
	assert.Equal(t, bad.Value(), actual.Value())

	top, err  =  h.Top()
	assert.NoError(t, err)
	actual, err = h.RemoveTop()
	assert.NoError(t, err)
	assert.Equal(t, actual.Value(), top.Value())
	assert.Equal(t, worse.Value(), actual.Value())

	_, err = h.Top()
	assert.Error(t, err)
	_, err = h.RemoveTop()
	assert.Error(t, err)


}
