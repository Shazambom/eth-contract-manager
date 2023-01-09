package storage

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)



func TestNewRedisWriter(t *testing.T) {
	writer := NewRedisWriter(rdsCfg)
	defer writer.Close()
	assert.IsType(t, &Redis{}, writer)
	inter := reflect.TypeOf((*RedisWriter)(nil)).Elem()
	assert.True(t, reflect.TypeOf(writer).Implements(inter))
}

func TestNewRedisListener(t *testing.T) {
	writer := NewRedisListener(rdsCfg)
	defer writer.Close()
	assert.IsType(t, &Redis{}, writer)
	inter := reflect.TypeOf((*RedisListener)(nil)).Elem()
	assert.True(t, reflect.TypeOf(writer).Implements(inter))
}

func TestRedis_VerifyValidAddress(t *testing.T) {
	writer := NewRedisWriter(rdsCfg)
	defer writer.Close()
	err := writer.VerifyValidAddress(ctx, "abc", "def")
	assert.Nil(t, err)
}

func TestRedis_VerifyValidAddress_RegisteredMultipleTimes(t *testing.T) {
	token, err := NewToken(
		"0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0",
		"0xVerifyValidAddress_RegisteredMultipleTimes",
		"0xce11e286abab09c3ad05f1f9fff4daaf4f5139214a1f4746661018f0b855f075",
		[]byte{127, 136, 18, 86, 140, 150, 133, 28, 2, 231, 67, 205, 183, 56, 83, 131, 117, 198, 87, 238, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 65, 191, 113, 122, 245, 174, 209, 169, 86, 54, 132, 170, 142, 180, 60, 227, 79, 228, 48, 47, 200, 218, 29, 40, 41, 84, 8, 6, 215, 214, 174, 233, 24, 86, 47, 76, 95, 125, 66, 14, 28, 105, 3, 155, 152, 238, 155, 181, 105, 94, 160, 122, 82, 231, 28, 209, 127, 202, 133, 33, 192, 186, 59, 250, 154, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		"450000000000000000")
	assert.Nil(t, err)
	writer := NewRedisWriter(rdsCfg)
	defer writer.Close()

	assert.Nil(t, writer.MarkAddressAsUsed(ctx, token))
	assert.Error(t, writer.VerifyValidAddress(ctx, token.UserAddress, token.ContractAddress))
}

func TestRedis_IncrementCounter(t *testing.T) {
	writer := NewRedisWriter(rdsCfg)
	defer writer.Close()
	assert.Nil(t, writer.IncrementCounter(ctx, 3, 1000, "0xIncrementCounter"))
	assert.Error(t, writer.IncrementCounter(ctx, 1000, 10, "0xIncrementCounter"))
}

func TestRedis_Get(t *testing.T) {
	token, err := NewToken(
		"0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0",
		"0xGet",
		"0xce11e286abab09c3ad05f1f9fff4daaf4f5139214a1f4746661018f0b855f075",
		[]byte{127, 136, 18, 86, 140, 150, 133, 28, 2, 231, 67, 205, 183, 56, 83, 131, 117, 198, 87, 238, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 65, 191, 113, 122, 245, 174, 209, 169, 86, 54, 132, 170, 142, 180, 60, 227, 79, 228, 48, 47, 200, 218, 29, 40, 41, 84, 8, 6, 215, 214, 174, 233, 24, 86, 47, 76, 95, 125, 66, 14, 28, 105, 3, 155, 152, 238, 155, 181, 105, 94, 160, 122, 82, 231, 28, 209, 127, 202, 133, 33, 192, 186, 59, 250, 154, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		"450000000000000000")
	assert.Nil(t, err)
	writer := NewRedisWriter(rdsCfg)
	defer writer.Close()

	assert.Nil(t, writer.MarkAddressAsUsed(ctx, token))
	tkn, getErr := writer.Get(ctx, token.UserAddress, token.ContractAddress)
	assert.Nil(t, getErr)
	assert.Equal(t, token, tkn)
}

func TestRedis_GetReservedCount(t *testing.T) {
	writer := NewRedisWriter(rdsCfg)
	defer writer.Close()
	assert.Nil(t, writer.GetReservedCount(ctx, 3, 1000, "0xGetReservedCount"))
}

func TestRedis_GetReservedCount_NumRequestedOver(t *testing.T) {
	writer := NewRedisWriter(rdsCfg)
	defer writer.Close()
	assert.Nil(t, writer.IncrementCounter(ctx, 1000, 1000, "0xTestRedis_GetReservedCount_NumRequestedOver"))
	assert.Error(t, writer.GetReservedCount(ctx, 3, 1000, "0xTestRedis_GetReservedCount_NumRequestedOver"))
}


func TestRedis_Listen(t *testing.T) {
	token, err := NewToken(
		"0xListen",
		"0xListen",
		"0xce11e286abab09c3ad05f1f9fff4daaf4f5139214a1f4746661018f0b855f075",
		[]byte{127, 136, 18, 86, 140, 150, 133, 28, 2, 231, 67, 205, 183, 56, 83, 131, 117, 198, 87, 238, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 65, 191, 113, 122, 245, 174, 209, 169, 86, 54, 132, 170, 142, 180, 60, 227, 79, 228, 48, 47, 200, 218, 29, 40, 41, 84, 8, 6, 215, 214, 174, 233, 24, 86, 47, 76, 95, 125, 66, 14, 28, 105, 3, 155, 152, 238, 155, 181, 105, 94, 160, 122, 82, 231, 28, 209, 127, 202, 133, 33, 192, 186, 59, 250, 154, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	"450000000000000000")
	assert.Nil(t, err)
	writer := NewRedisWriter(rdsCfg)
	defer writer.Close()
	listener := NewRedisListener(rdsCfg)
	defer listener.Close()
	assert.Nil(t, listener.InitEvents())
	finishVal := "done :)"
	done := make(chan string)
	go listener.Listen(func(key string, val string, err error) error {
		if key != token.ContractAddress + "_" + token.UserAddress {
			return nil
		}
		assert.Nil(t, err)
		tkn := Token{}
		assert.Nil(t, json.Unmarshal([]byte(val), &tkn))
		assert.Equal(t, token, &tkn)
		done <- finishVal
		return errors.New("stop")
	})
	time.Sleep(1 *time.Second)
	assert.Nil(t, writer.MarkAddressAsUsed(ctx, token))
	assert.Equal(t, finishVal, <-done)
}






