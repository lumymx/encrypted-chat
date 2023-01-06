package main

import (
	"crypto/aes"
	"crypto/cipher"
)

func EncryptMessage(message, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	encryptedMessage := make([]byte, len(message))
	stream.XORKeyStream(encryptedMessage, message)
	return encryptedMessage, nil
}

func DecryptMessage(encryptedMessage, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	stream := cipher.NewCFBDecrypter(block, iv)
	decryptedMessage := make([]byte, len(encryptedMessage))
	stream.XORKeyStream(decryptedMessage, encryptedMessage)
	return encryptedMessage, nil
}
