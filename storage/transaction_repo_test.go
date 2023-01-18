package storage

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNewTransactionRepo(t *testing.T) {
	assert.IsType(t, &TransactionRepo{}, tr)
	inter := reflect.TypeOf((*TransactionRepository)(nil)).Elem()
	assert.True(t, reflect.TypeOf(tr).Implements(inter))
}

func TestTransactionRepo_StoreTransaction(t *testing.T) {
	token, tokenErr := NewToken(
		"0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0",
		"0x0fA37C622C7E57A06ba12afF48c846F42241F7F0",
		"0xce11e286abab09c3ad05f1f9fff4daaf4f5139214a1f4746661018f0b855f075",
		[]byte{127, 136, 18, 86, 140, 150, 133, 28, 2, 231, 67, 205, 183, 56, 83, 131, 117, 198, 87, 238, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 65, 191, 113, 122, 245, 174, 209, 169, 86, 54, 132, 170, 142, 180, 60, 227, 79, 228, 48, 47, 200, 218, 29, 40, 41, 84, 8, 6, 215, 214, 174, 233, 24, 86, 47, 76, 95, 125, 66, 14, 28, 105, 3, 155, 152, 238, 155, 181, 105, 94, 160, 122, 82, 231, 28, 209, 127, 202, 133, 33, 192, 186, 59, 250, 154, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		"450000000000000000")
	assert.Nil(t, tokenErr)
	assert.Nil(t, tr.StoreTransaction(ctx, *token))
}

func TestTransactionRepo_DeleteTransaction(t *testing.T) {
	token, tokenErr := NewToken(
		"0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0",
		"0x0fA37C622C7E57A06ba12afF48c846F42241F7F0",
		"0xce11e286abab09c3ad05f1f9fff4daaf4f5139214a1f4746661018f0b855f075",
		[]byte{127, 136, 18, 86, 140, 150, 133, 28, 2, 231, 67, 205, 183, 56, 83, 131, 117, 198, 87, 238, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 65, 191, 113, 122, 245, 174, 209, 169, 86, 54, 132, 170, 142, 180, 60, 227, 79, 228, 48, 47, 200, 218, 29, 40, 41, 84, 8, 6, 215, 214, 174, 233, 24, 86, 47, 76, 95, 125, 66, 14, 28, 105, 3, 155, 152, 238, 155, 181, 105, 94, 160, 122, 82, 231, 28, 209, 127, 202, 133, 33, 192, 186, 59, 250, 154, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		"450000000000000000")
	assert.Nil(t, tokenErr)
	assert.Nil(t, tr.StoreTransaction(ctx, *token))
	assert.Nil(t, tr.DeleteTransaction(ctx, "0x0fA37C622C7E57A06ba12afF48c846F42241F7F0", "0xce11e286abab09c3ad05f1f9fff4daaf4f5139214a1f4746661018f0b855f075"))
}

func TestTransactionRepo_GetTransactions(t *testing.T) {
	template, templateErr := NewToken(
		"0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0",
		"TestTransactionRepo_GetTransactions",
		"",
		[]byte{127, 136, 18, 86, 140, 150, 133, 28, 2, 231, 67, 205, 183, 56, 83, 131, 117, 198, 87, 238, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 65, 191, 113, 122, 245, 174, 209, 169, 86, 54, 132, 170, 142, 180, 60, 227, 79, 228, 48, 47, 200, 218, 29, 40, 41, 84, 8, 6, 215, 214, 174, 233, 24, 86, 47, 76, 95, 125, 66, 14, 28, 105, 3, 155, 152, 238, 155, 181, 105, 94, 160, 122, 82, 231, 28, 209, 127, 202, 133, 33, 192, 186, 59, 250, 154, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		"450000000000000000")
	assert.Nil(t, templateErr)

	tokenHash1, hash1Err := NewToken(
		"0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0",
		"TestTransactionRepo_GetTransactions",
		"1",
		[]byte{127, 136, 18, 86, 140, 150, 133, 28, 2, 231, 67, 205, 183, 56, 83, 131, 117, 198, 87, 238, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 65, 191, 113, 122, 245, 174, 209, 169, 86, 54, 132, 170, 142, 180, 60, 227, 79, 228, 48, 47, 200, 218, 29, 40, 41, 84, 8, 6, 215, 214, 174, 233, 24, 86, 47, 76, 95, 125, 66, 14, 28, 105, 3, 155, 152, 238, 155, 181, 105, 94, 160, 122, 82, 231, 28, 209, 127, 202, 133, 33, 192, 186, 59, 250, 154, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		"450000000000000000")
	assert.Nil(t, hash1Err)
	assert.Nil(t, tr.StoreTransaction(ctx, *tokenHash1))

	tokenHash2, hash2Err := NewToken(
		"0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0",
		"TestTransactionRepo_GetTransactions",
		"2",
		[]byte{127, 136, 18, 86, 140, 150, 133, 28, 2, 231, 67, 205, 183, 56, 83, 131, 117, 198, 87, 238, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 65, 191, 113, 122, 245, 174, 209, 169, 86, 54, 132, 170, 142, 180, 60, 227, 79, 228, 48, 47, 200, 218, 29, 40, 41, 84, 8, 6, 215, 214, 174, 233, 24, 86, 47, 76, 95, 125, 66, 14, 28, 105, 3, 155, 152, 238, 155, 181, 105, 94, 160, 122, 82, 231, 28, 209, 127, 202, 133, 33, 192, 186, 59, 250, 154, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		"450000000000000000")
	assert.Nil(t, hash2Err)
	assert.Nil(t, tr.StoreTransaction(ctx, *tokenHash2))
	tokenHash3, hash3Err := NewToken(
		"0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0",
		"TestTransactionRepo_GetTransactions",
		"3",
		[]byte{127, 136, 18, 86, 140, 150, 133, 28, 2, 231, 67, 205, 183, 56, 83, 131, 117, 198, 87, 238, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 65, 191, 113, 122, 245, 174, 209, 169, 86, 54, 132, 170, 142, 180, 60, 227, 79, 228, 48, 47, 200, 218, 29, 40, 41, 84, 8, 6, 215, 214, 174, 233, 24, 86, 47, 76, 95, 125, 66, 14, 28, 105, 3, 155, 152, 238, 155, 181, 105, 94, 160, 122, 82, 231, 28, 209, 127, 202, 133, 33, 192, 186, 59, 250, 154, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		"450000000000000000")
	assert.Nil(t, hash3Err)
	assert.Nil(t, tr.StoreTransaction(ctx, *tokenHash3))
	tokenHash4, hash4Err := NewToken(
		"0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0",
		"TestTransactionRepo_GetTransactions",
		"4",
		[]byte{127, 136, 18, 86, 140, 150, 133, 28, 2, 231, 67, 205, 183, 56, 83, 131, 117, 198, 87, 238, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 65, 191, 113, 122, 245, 174, 209, 169, 86, 54, 132, 170, 142, 180, 60, 227, 79, 228, 48, 47, 200, 218, 29, 40, 41, 84, 8, 6, 215, 214, 174, 233, 24, 86, 47, 76, 95, 125, 66, 14, 28, 105, 3, 155, 152, 238, 155, 181, 105, 94, 160, 122, 82, 231, 28, 209, 127, 202, 133, 33, 192, 186, 59, 250, 154, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		"450000000000000000")
	assert.Nil(t, hash4Err)
	assert.Nil(t, tr.StoreTransaction(ctx, *tokenHash4))
	tokens, err := tr.GetTransactions(ctx, "TestTransactionRepo_GetTransactions")
	assert.Nil(t, err)
	assert.Equal(t, 4, len(tokens))
	for i := 0; i < 4; i++ {
		found := false
		template.Hash = fmt.Sprintf("%d", 1+i)
		for _, token := range tokens {
			if token.ContractAddress == template.ContractAddress &&
				token.UserAddress == template.UserAddress &&
				token.Value == template.Value &&
				token.Hash == template.Hash &&
				string(token.ABIPackedTxn) == string(template.ABIPackedTxn) &&
				token.IsComplete == template.IsComplete {
				found = true
			}
		}
		assert.True(t, found)
	}

	assert.Nil(t, tr.DeleteTransaction(ctx, "TestTransactionRepo_GetTransactions", "4"))

	tokensNext, errNext := tr.GetTransactions(ctx, "TestTransactionRepo_GetTransactions")
	assert.Nil(t, errNext)
	assert.Equal(t, 3, len(tokensNext))
	for i := 0; i < 3; i++ {
		found := false
		template.Hash = fmt.Sprintf("%d", 1+i)
		for _, token := range tokensNext {
			if token.ContractAddress == template.ContractAddress &&
				token.UserAddress == template.UserAddress &&
				token.Value == template.Value &&
				token.Hash == template.Hash &&
				string(token.ABIPackedTxn) == string(template.ABIPackedTxn) &&
				token.IsComplete == template.IsComplete{
				found = true
			}
		}
		assert.True(t, found)
	}
}

func TestTransactionRepo_CompleteTransaction(t *testing.T) {
	template, templateErr := NewToken(
		"0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0",
		"TestTransactionRepo_CompleteTransaction",
		"",
		[]byte{127, 136, 18, 86, 140, 150, 133, 28, 2, 231, 67, 205, 183, 56, 83, 131, 117, 198, 87, 238, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 65, 191, 113, 122, 245, 174, 209, 169, 86, 54, 132, 170, 142, 180, 60, 227, 79, 228, 48, 47, 200, 218, 29, 40, 41, 84, 8, 6, 215, 214, 174, 233, 24, 86, 47, 76, 95, 125, 66, 14, 28, 105, 3, 155, 152, 238, 155, 181, 105, 94, 160, 122, 82, 231, 28, 209, 127, 202, 133, 33, 192, 186, 59, 250, 154, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		"450000000000000000")
	assert.Nil(t, templateErr)

	tokenHash1, hash1Err := NewToken(
		"0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0",
		"TestTransactionRepo_CompleteTransaction",
		"1",
		[]byte{127, 136, 18, 86, 140, 150, 133, 28, 2, 231, 67, 205, 183, 56, 83, 131, 117, 198, 87, 238, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 65, 191, 113, 122, 245, 174, 209, 169, 86, 54, 132, 170, 142, 180, 60, 227, 79, 228, 48, 47, 200, 218, 29, 40, 41, 84, 8, 6, 215, 214, 174, 233, 24, 86, 47, 76, 95, 125, 66, 14, 28, 105, 3, 155, 152, 238, 155, 181, 105, 94, 160, 122, 82, 231, 28, 209, 127, 202, 133, 33, 192, 186, 59, 250, 154, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		"450000000000000000")
	assert.Nil(t, hash1Err)
	assert.Nil(t, tr.StoreTransaction(ctx, *tokenHash1))

	tokenHash2, hash2Err := NewToken(
		"0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0",
		"TestTransactionRepo_CompleteTransaction",
		"2",
		[]byte{127, 136, 18, 86, 140, 150, 133, 28, 2, 231, 67, 205, 183, 56, 83, 131, 117, 198, 87, 238, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 65, 191, 113, 122, 245, 174, 209, 169, 86, 54, 132, 170, 142, 180, 60, 227, 79, 228, 48, 47, 200, 218, 29, 40, 41, 84, 8, 6, 215, 214, 174, 233, 24, 86, 47, 76, 95, 125, 66, 14, 28, 105, 3, 155, 152, 238, 155, 181, 105, 94, 160, 122, 82, 231, 28, 209, 127, 202, 133, 33, 192, 186, 59, 250, 154, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		"450000000000000000")
	assert.Nil(t, hash2Err)
	assert.Nil(t, tr.StoreTransaction(ctx, *tokenHash2))
	tokenHash3, hash3Err := NewToken(
		"0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0",
		"TestTransactionRepo_CompleteTransaction",
		"3",
		[]byte{127, 136, 18, 86, 140, 150, 133, 28, 2, 231, 67, 205, 183, 56, 83, 131, 117, 198, 87, 238, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 65, 191, 113, 122, 245, 174, 209, 169, 86, 54, 132, 170, 142, 180, 60, 227, 79, 228, 48, 47, 200, 218, 29, 40, 41, 84, 8, 6, 215, 214, 174, 233, 24, 86, 47, 76, 95, 125, 66, 14, 28, 105, 3, 155, 152, 238, 155, 181, 105, 94, 160, 122, 82, 231, 28, 209, 127, 202, 133, 33, 192, 186, 59, 250, 154, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		"450000000000000000")
	assert.Nil(t, hash3Err)
	assert.Nil(t, tr.StoreTransaction(ctx, *tokenHash3))
	tokenHash4, hash4Err := NewToken(
		"0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0",
		"TestTransactionRepo_CompleteTransaction",
		"4",
		[]byte{127, 136, 18, 86, 140, 150, 133, 28, 2, 231, 67, 205, 183, 56, 83, 131, 117, 198, 87, 238, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 65, 191, 113, 122, 245, 174, 209, 169, 86, 54, 132, 170, 142, 180, 60, 227, 79, 228, 48, 47, 200, 218, 29, 40, 41, 84, 8, 6, 215, 214, 174, 233, 24, 86, 47, 76, 95, 125, 66, 14, 28, 105, 3, 155, 152, 238, 155, 181, 105, 94, 160, 122, 82, 231, 28, 209, 127, 202, 133, 33, 192, 186, 59, 250, 154, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		"450000000000000000")
	assert.Nil(t, hash4Err)
	assert.Nil(t, tr.StoreTransaction(ctx, *tokenHash4))

	tokens, err := tr.GetTransactions(ctx, "TestTransactionRepo_CompleteTransaction")
	assert.Nil(t, err)
	assert.Equal(t, 4, len(tokens))
	for i := 0; i < 4; i++ {
		found := false
		template.Hash = fmt.Sprintf("%d", 1+i)
		for _, token := range tokens {
			if token.ContractAddress == template.ContractAddress &&
				token.UserAddress == template.UserAddress &&
				token.Value == template.Value &&
				token.Hash == template.Hash &&
				string(token.ABIPackedTxn) == string(template.ABIPackedTxn) &&
				token.IsComplete == template.IsComplete{
				found = true
			}
		}
		assert.True(t, found)
	}

	assert.Nil(t, tr.CompleteTransaction(ctx, "TestTransactionRepo_CompleteTransaction", "4"))

	tokensNext, errNext := tr.GetTransactions(ctx, "TestTransactionRepo_CompleteTransaction")
	assert.Nil(t, errNext)
	assert.Equal(t, 3, len(tokensNext))
	for i := 0; i < 3; i++ {
		found := false
		template.Hash = fmt.Sprintf("%d", 1+i)
		for _, token := range tokensNext {
			if token.ContractAddress == template.ContractAddress &&
				token.UserAddress == template.UserAddress &&
				token.Value == template.Value &&
				token.Hash == template.Hash &&
				string(token.ABIPackedTxn) == string(template.ABIPackedTxn) &&
				token.IsComplete == template.IsComplete {
				found = true
			}
		}
		assert.True(t, found)
	}
}

func TestTransactionRepo_GetAllTransactions(t *testing.T) {
	template, templateErr := NewToken(
		"0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0",
		"TestTransactionRepo_GetAllTransactions",
		"",
		[]byte{127, 136, 18, 86, 140, 150, 133, 28, 2, 231, 67, 205, 183, 56, 83, 131, 117, 198, 87, 238, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 65, 191, 113, 122, 245, 174, 209, 169, 86, 54, 132, 170, 142, 180, 60, 227, 79, 228, 48, 47, 200, 218, 29, 40, 41, 84, 8, 6, 215, 214, 174, 233, 24, 86, 47, 76, 95, 125, 66, 14, 28, 105, 3, 155, 152, 238, 155, 181, 105, 94, 160, 122, 82, 231, 28, 209, 127, 202, 133, 33, 192, 186, 59, 250, 154, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		"450000000000000000")
	assert.Nil(t, templateErr)

	tokenHash1, hash1Err := NewToken(
		"0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0",
		"TestTransactionRepo_GetAllTransactions",
		"1",
		[]byte{127, 136, 18, 86, 140, 150, 133, 28, 2, 231, 67, 205, 183, 56, 83, 131, 117, 198, 87, 238, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 65, 191, 113, 122, 245, 174, 209, 169, 86, 54, 132, 170, 142, 180, 60, 227, 79, 228, 48, 47, 200, 218, 29, 40, 41, 84, 8, 6, 215, 214, 174, 233, 24, 86, 47, 76, 95, 125, 66, 14, 28, 105, 3, 155, 152, 238, 155, 181, 105, 94, 160, 122, 82, 231, 28, 209, 127, 202, 133, 33, 192, 186, 59, 250, 154, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		"450000000000000000")
	assert.Nil(t, hash1Err)
	assert.Nil(t, tr.StoreTransaction(ctx, *tokenHash1))

	tokenHash2, hash2Err := NewToken(
		"0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0",
		"TestTransactionRepo_GetAllTransactions",
		"2",
		[]byte{127, 136, 18, 86, 140, 150, 133, 28, 2, 231, 67, 205, 183, 56, 83, 131, 117, 198, 87, 238, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 65, 191, 113, 122, 245, 174, 209, 169, 86, 54, 132, 170, 142, 180, 60, 227, 79, 228, 48, 47, 200, 218, 29, 40, 41, 84, 8, 6, 215, 214, 174, 233, 24, 86, 47, 76, 95, 125, 66, 14, 28, 105, 3, 155, 152, 238, 155, 181, 105, 94, 160, 122, 82, 231, 28, 209, 127, 202, 133, 33, 192, 186, 59, 250, 154, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		"450000000000000000")
	assert.Nil(t, hash2Err)
	assert.Nil(t, tr.StoreTransaction(ctx, *tokenHash2))
	tokenHash3, hash3Err := NewToken(
		"0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0",
		"TestTransactionRepo_GetAllTransactions",
		"3",
		[]byte{127, 136, 18, 86, 140, 150, 133, 28, 2, 231, 67, 205, 183, 56, 83, 131, 117, 198, 87, 238, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 65, 191, 113, 122, 245, 174, 209, 169, 86, 54, 132, 170, 142, 180, 60, 227, 79, 228, 48, 47, 200, 218, 29, 40, 41, 84, 8, 6, 215, 214, 174, 233, 24, 86, 47, 76, 95, 125, 66, 14, 28, 105, 3, 155, 152, 238, 155, 181, 105, 94, 160, 122, 82, 231, 28, 209, 127, 202, 133, 33, 192, 186, 59, 250, 154, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		"450000000000000000")
	assert.Nil(t, hash3Err)
	assert.Nil(t, tr.StoreTransaction(ctx, *tokenHash3))
	tokenHash4, hash4Err := NewToken(
		"0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0",
		"TestTransactionRepo_GetAllTransactions",
		"4",
		[]byte{127, 136, 18, 86, 140, 150, 133, 28, 2, 231, 67, 205, 183, 56, 83, 131, 117, 198, 87, 238, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 65, 191, 113, 122, 245, 174, 209, 169, 86, 54, 132, 170, 142, 180, 60, 227, 79, 228, 48, 47, 200, 218, 29, 40, 41, 84, 8, 6, 215, 214, 174, 233, 24, 86, 47, 76, 95, 125, 66, 14, 28, 105, 3, 155, 152, 238, 155, 181, 105, 94, 160, 122, 82, 231, 28, 209, 127, 202, 133, 33, 192, 186, 59, 250, 154, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		"450000000000000000")
	assert.Nil(t, hash4Err)
	assert.Nil(t, tr.StoreTransaction(ctx, *tokenHash4))

	tokens, err := tr.GetAllTransactions(ctx, "TestTransactionRepo_GetAllTransactions")
	assert.Nil(t, err)
	assert.Equal(t, 4, len(tokens))
	for i := 0; i < 4; i++ {
		found := false
		template.Hash = fmt.Sprintf("%d", 1+i)
		for _, token := range tokens {
			if token.ContractAddress == template.ContractAddress &&
				token.UserAddress == template.UserAddress &&
				token.Value == template.Value &&
				token.Hash == template.Hash &&
				string(token.ABIPackedTxn) == string(template.ABIPackedTxn) &&
				token.IsComplete == template.IsComplete {
				found = true
			}
		}
		assert.True(t, found)
	}

	assert.Nil(t, tr.CompleteTransaction(ctx, "TestTransactionRepo_GetAllTransactions", "4"))

	tokensNext, errNext := tr.GetAllTransactions(ctx, "TestTransactionRepo_GetAllTransactions")
	assert.Nil(t, errNext)
	assert.Equal(t, 4, len(tokensNext))
	for i := 0; i < 4; i++ {
		found := false
		template.Hash = fmt.Sprintf("%d", 1+i)
		if i == 3 {
			template.IsComplete = true
		}
		for _, token := range tokensNext {
			if token.ContractAddress == template.ContractAddress &&
				token.UserAddress == template.UserAddress &&
				token.Value == template.Value &&
				token.Hash == template.Hash &&
				string(token.ABIPackedTxn) == string(template.ABIPackedTxn) &&
				token.IsComplete == template.IsComplete {
				found = true
			}
		}
		assert.True(t, found)
	}
}

