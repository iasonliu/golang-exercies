package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"io"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	msg := "Tohis is totally fun get hands"

	password := "ilovedogs"
	// using 16 bit for password bcrypt as a Key for aes.NewCipher
	key, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	// warp io.writer
	wtr := &bytes.Buffer{}
	encWriter, err := encryptWriter(wtr, key[:16])
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = io.WriteString(encWriter, msg)
	if err != nil {
		log.Fatalln(err.Error())
	}

	encrypted := wtr.String()
	fmt.Println("before beae64, requier base64:", encrypted)
	// other way
	rsEncode, err := enDecode(key[:16], msg)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("before beae64, requier base64:", string(rsEncode))
	fmt.Println("ASE encryption:", base64.URLEncoding.EncodeToString(rsEncode))

	// Decode
	rsDecode, err := enDecode(key[:16], string(rsEncode))
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("Decode for ASE: ", string(rsDecode))
}

func enDecode(key []byte, msg string) ([]byte, error) {
	b, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("coludn't newCipther %w", err)
	}
	// initalization vector
	iv := make([]byte, aes.BlockSize)

	s := cipher.NewCTR(b, iv)

	buff := &bytes.Buffer{}
	sw := cipher.StreamWriter{
		S: s,
		W: buff,
	}
	_, err = sw.Write([]byte(msg))
	if err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}

func encryptWriter(wtr io.Writer, key []byte) (io.Writer, error) {
	b, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("coludn't newCipther %w", err)
	}
	// initalization vector
	iv := make([]byte, aes.BlockSize)

	s := cipher.NewCTR(b, iv)

	return cipher.StreamWriter{
		S: s,
		W: wtr,
	}, nil
}
