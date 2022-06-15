package storage

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/stretchr/testify/assert"
	"path"
	"testing"
)

var cfg = &aws.Config{
	Endpoint: aws.String("localhost:4566"),
	Region: aws.String("us-east-1"),
	Credentials: credentials.NewStaticCredentials("xxx", "yyy", ""),
	S3ForcePathStyle: aws.Bool(true),
	DisableSSL: aws.Bool(true),
}
var testBucketName = "bucket"


func TestNewS3(t *testing.T) {
	s3, err := NewS3(cfg, testBucketName)
	assert.Nil(t, err)
	assert.IsType(t, &S3{}, s3)
}

func TestS3_InitBucket(t *testing.T) {
	s3, err := NewS3(cfg, testBucketName)
	assert.Nil(t, err)
	assert.Nil(t, s3.InitBucket())
}

func TestS3_StoreToken(t *testing.T) {
	s3, err := NewS3(cfg, testBucketName)
	assert.Nil(t, err)
	assert.Nil(t, s3.InitBucket())
	token := NewToken(
		"0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0",
		"0x0fA37C622C7E57A06ba12afF48c846F42241F7F0",
		"0xce11e286abab09c3ad05f1f9fff4daaf4f5139214a1f4746661018f0b855f075",
		"abc",
		[]byte{127, 136, 18, 86, 140, 150, 133, 28, 2, 231, 67, 205, 183, 56, 83, 131, 117, 198, 87, 238, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 65, 191, 113, 122, 245, 174, 209, 169, 86, 54, 132, 170, 142, 180, 60, 227, 79, 228, 48, 47, 200, 218, 29, 40, 41, 84, 8, 6, 215, 214, 174, 233, 24, 86, 47, 76, 95, 125, 66, 14, 28, 105, 3, 155, 152, 238, 155, 181, 105, 94, 160, 122, 82, 231, 28, 209, 127, 202, 133, 33, 192, 186, 59, 250, 154, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		3)
	storeErr := s3.StoreToken(token)
	assert.Nil(t, storeErr)
}

func TestS3_GetToken(t *testing.T) {
	s3, err := NewS3(cfg, testBucketName)
	assert.Nil(t, err)
	assert.Nil(t, s3.InitBucket())
	token := NewToken(
		"0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0",
		"0x0fA37C622C7E57A06ba12afF48c846F42241F7F0",
		"0xce11e286abab09c3ad05f1f9fff4daaf4f5139214a1f4746661018f0b855f075",
		"abc",
		[]byte{127, 136, 18, 86, 140, 150, 133, 28, 2, 231, 67, 205, 183, 56, 83, 131, 117, 198, 87, 238, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 65, 191, 113, 122, 245, 174, 209, 169, 86, 54, 132, 170, 142, 180, 60, 227, 79, 228, 48, 47, 200, 218, 29, 40, 41, 84, 8, 6, 215, 214, 174, 233, 24, 86, 47, 76, 95, 125, 66, 14, 28, 105, 3, 155, 152, 238, 155, 181, 105, 94, 160, 122, 82, 231, 28, 209, 127, 202, 133, 33, 192, 186, 59, 250, 154, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		3)
	storeErr := s3.StoreToken(token)
	assert.Nil(t, storeErr)

	retrievedToken, getErr := s3.GetToken(token.ContractAddress, token.UserAddress)
	assert.Nil(t, getErr)
	assert.Equal(t, token, retrievedToken)
}


func TestS3_ListKeys(t *testing.T) {
	s3, err := NewS3(cfg, testBucketName)
	assert.Nil(t, err)
	assert.Nil(t, s3.InitBucket())
	token := NewToken(
		"0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0",
		"0x0fA37C622C7E57A06ba12afF48c846F42241F7F0",
		"0xce11e286abab09c3ad05f1f9fff4daaf4f5139214a1f4746661018f0b855f075",
		"abc",
		[]byte{127, 136, 18, 86, 140, 150, 133, 28, 2, 231, 67, 205, 183, 56, 83, 131, 117, 198, 87, 238, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 65, 191, 113, 122, 245, 174, 209, 169, 86, 54, 132, 170, 142, 180, 60, 227, 79, 228, 48, 47, 200, 218, 29, 40, 41, 84, 8, 6, 215, 214, 174, 233, 24, 86, 47, 76, 95, 125, 66, 14, 28, 105, 3, 155, 152, 238, 155, 181, 105, 94, 160, 122, 82, 231, 28, 209, 127, 202, 133, 33, 192, 186, 59, 250, 154, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		3)
	storeErr := s3.StoreToken(token)
	assert.Nil(t, storeErr)

	keys, keyErr := s3.ListKeys()
	assert.Nil(t, keyErr)
	assert.Equal(t, path.Join(token.ContractAddress, token.UserAddress), keys[0])
}
