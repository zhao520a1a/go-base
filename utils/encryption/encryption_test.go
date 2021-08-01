package encryption

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func prepare() {
	SetConfig(map[int]string{
		1: "97ea2c7a8533af02050991b363e8c97f",
	})
}

func TestEncryptText(t *testing.T) {
	prepare()

	ctx := context.Background()
	text := "13488660468"

	cipherText, err := EncryptText(ctx, text)
	assert.NoError(t, err)
	log.Println(cipherText)
}

func TestDecryptText(t *testing.T) {
	prepare()

	ctx := context.Background()
	cipherText := "CAESEDEb+FJGLYfCbPPqGoLbNPE="

	text, err := DecryptText(ctx, cipherText)
	assert.NoError(t, err)
	log.Println(text)
}
