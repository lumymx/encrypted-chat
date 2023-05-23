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
	encryptedMessage := make([]byte, len(message))
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encryptedMessage, message)
	return encryptedMessage, nil
}

func DecryptMessage(encryptedMessage, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	decryptedMessage := make([]byte, len(encryptedMessage))
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(decryptedMessage, encryptedMessage)
	return decryptedMessage, nil
}
