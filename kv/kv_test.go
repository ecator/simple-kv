package kv_test

import (
	"simple-kv/kv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetValue(t *testing.T) {
	token := "token1"
	key := "key1"
	value := "value1"
	assert.EqualValues(t, "", kv.GetValue(token, key))
	assert.EqualValues(t, nil, kv.SetValue(token, key, value))
	assert.EqualValues(t, value, kv.GetValue(token, key))
	value = "vaule2"
	assert.EqualValues(t, nil, kv.SetValue(token, key, value))
	assert.EqualValues(t, value, kv.GetValue(token, key))
}
