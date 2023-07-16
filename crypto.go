package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func GenerateKeyAndIV() ([]byte, []byte, error) {
	key := make([]byte, 16)
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return nil, nil, err
	}
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, nil, err
	}
	return key, iv, nil
}

func EncryptMessage(message, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	encryptedMessage := make([]byte, aes.BlockSize+len(message))
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(encryptedMessage[aes.BlockSize:], message)
	copy(encryptedMessage[:aes.BlockSize], iv)
	return encryptedMessage, nil
}

func DecryptMessage(encryptedMessage, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(encryptedMessage) < aes.BlockSize {
		return nil, fmt.Errorf("encrypted message too short")
	}
	iv = encryptedMessage[:aes.BlockSize]
	encryptedMessage = encryptedMessage[aes.BlockSize:]
	decryptedMessage := make([]byte, len(encryptedMessage))
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(decryptedMessage, encryptedMessage)
	return decryptedMessage, nil
}
