package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
)

func AesGCMEncrypt(plaintext []byte) (nonce, ciphertext, aesKey []byte, err error) {
	aesKey = make([]byte, 32)
	_, err = rand.Read(aesKey)
	if err != nil {
		return
	}

	nonce = make([]byte, 12)
	_, err = rand.Read(nonce)
	if err != nil {
		return
	}

	aesBlockCipher, err := aes.NewCipher(aesKey)
	if err != nil {
		return
	}

	aead, err := cipher.NewGCMWithNonceSize(aesBlockCipher, 12)
	ciphertext = aead.Seal(nil, nonce, plaintext, nil)

	return
}

func AesGCMDecrypt(nonce, ciphertext, key []byte) (plaintext []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		err = fmt.Errorf("failed to new cipher (%w)", err)
		return
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		err = fmt.Errorf("failed to new GCM (%w)", err)
		return
	}

	plaintext, err = aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		err = fmt.Errorf("failed to decrypt (%w)", err)
		return
	}

	return
}

func main() {
	nonce, cipher, key, err := AesGCMEncrypt([]byte("hello world"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("ciphertext: \n%v\n\n", cipher)

	plaintext, err := AesGCMDecrypt(nonce, cipher, key)
	if err != nil {
		panic(err)
	}

	fmt.Printf("plaintext: \n%v\n", string(plaintext))
}
