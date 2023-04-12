package storage

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNewPrivateKeyRepository(t *testing.T) {
	assert.IsType(t, &PrivateKeyRepo{}, pkr)
	inter := reflect.TypeOf((*PrivateKeyRepository)(nil)).Elem()
	assert.True(t, reflect.TypeOf(pkr).Implements(inter))
}

func TestPrivateKeyRepo_UpsertPrivateKey(t *testing.T) {
	upsertErr := pkr.UpsertPrivateKey(ctx, "0xUpsertPrivateKey", "PrivateKeyTest")
	assert.Nil(t, upsertErr)
}

func TestPrivateKeyRepo_GetPrivateKey(t *testing.T) {
	contractAddress := "0xGetPrivateKey"
	key := "GetPrivateKeyTest"

	upsertErr := pkr.UpsertPrivateKey(ctx, contractAddress, key)
	assert.Nil(t, upsertErr)

	retrievedKey, getErr := pkr.GetPrivateKey(ctx, contractAddress)
	assert.Nil(t, getErr)
	assert.Equal(t, key, retrievedKey)
}

func TestPrivateKeyRepo_DeletePrivateKey(t *testing.T) {
	contractAddress := "0xDeletePrivateKey"
	key := "DeletePrivateKeyTest"

	upsertErr := pkr.UpsertPrivateKey(ctx, contractAddress, key)
	assert.Nil(t, upsertErr)

	retrievedKey, getErr := pkr.GetPrivateKey(ctx, contractAddress)
	assert.Nil(t, getErr)
	assert.Equal(t, key, retrievedKey)

	delErr := pkr.DeletePrivateKey(ctx, contractAddress)
	assert.Nil(t, delErr)

	_, afterDelGetErr := pkr.GetPrivateKey(ctx, contractAddress)
	assert.Error(t, afterDelGetErr)
	fmt.Println(afterDelGetErr.Error())
}
