package storage

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)



func TestNewPrivateKeyRepository(t *testing.T) {
	privateRepo, err := NewPrivateKeyRepository(TestPrivateKeyTableName, dynamoCfg)
	assert.Nil(t, err)
	assert.IsType(t, &PrivateKeyRepo{}, privateRepo)
	inter := reflect.TypeOf((*PrivateKeyRepository)(nil)).Elem()
	assert.True(t, reflect.TypeOf(privateRepo).Implements(inter))
}

func TestPrivateKeyRepo_UpsertPrivateKey(t *testing.T) {
	privateRepo, err := NewPrivateKeyRepository(TestPrivateKeyTableName, dynamoCfg)
	assert.Nil(t, err)

	upsertErr := privateRepo.UpsertPrivateKey(ctx, "0xUpsertPrivateKey", "PrivateKeyTest")
	assert.Nil(t, upsertErr)
}

func TestPrivateKeyRepo_GetPrivateKey(t *testing.T) {
	privateRepo, err := NewPrivateKeyRepository(TestPrivateKeyTableName, dynamoCfg)
	assert.Nil(t, err)

	contractAddress := "0xGetPrivateKey"
	key := "GetPrivateKeyTest"

	upsertErr := privateRepo.UpsertPrivateKey(ctx, contractAddress, key)
	assert.Nil(t, upsertErr)

	retrievedKey, getErr := privateRepo.GetPrivateKey(ctx, contractAddress)
	assert.Nil(t, getErr)
	assert.Equal(t, key, retrievedKey)
}

func TestPrivateKeyRepo_DeletePrivateKey(t *testing.T) {
	privateRepo, err := NewPrivateKeyRepository(TestPrivateKeyTableName, dynamoCfg)
	assert.Nil(t, err)

	contractAddress := "0xDeletePrivateKey"
	key := "DeletePrivateKeyTest"

	upsertErr := privateRepo.UpsertPrivateKey(ctx, contractAddress, key)
	assert.Nil(t, upsertErr)

	retrievedKey, getErr := privateRepo.GetPrivateKey(ctx, contractAddress)
	assert.Nil(t, getErr)
	assert.Equal(t, key, retrievedKey)

	delErr := privateRepo.DeletePrivateKey(ctx, contractAddress)
	assert.Nil(t, delErr)

	_, afterDelGetErr := privateRepo.GetPrivateKey(ctx, contractAddress)
	assert.Error(t, afterDelGetErr)
	fmt.Println(afterDelGetErr.Error())
}