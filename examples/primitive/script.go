package main

import "github.com/stretchr/testify/require"

func Setup() interface{} {
	return 1
}

func Default(assert *require.Assertions, data interface{}) {
	val, ok := data.(int64)

	assert.True(ok, "primitive data type OK")
	assert.Equal(1, int(val), "primitive data value OK")
}
