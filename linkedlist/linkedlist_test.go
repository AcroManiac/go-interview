package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	l := &LinkedList{}
	l.Add("A", 0)
	l.Add("D", 1)
	l.Add("C", 1)
	l.Add("B", 1)
	l.Add("Z", 0)

	assert.Equal(t, "Z, A, B, C, D", l.String())
}
