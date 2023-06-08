package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefault(t *testing.T) {
	a := assert.New(t)

	a.NotPanics(func() { Default(a) })
}
