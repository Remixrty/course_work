package Shifr

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	KEY = "FRNX"
	IV  = "23456789"
)

func Desc1(con *gin.Context) {
	//miwen := "7379533330307A63724F554661783369495A653253393533775377426948704A"
	logrus.Printf("All_OK")
	miwen, err := ioutil.ReadFile("DescStart.txt")
	if err != nil {
		fmt.Println(err)
	}
	str := decode(string(miwen))
	//fmt.Println(str)
	err = ioutil.WriteFile("DescEnd.data", []byte(str), 0777)
	if err != nil {
		fmt.Println(err)
	}
	con.JSON(http.StatusOK, gin.H{
		"str": str,
	})
}

func genKey(key string) string {
	if len(key) < 32 {
		key += strings.Repeat("0", 32)
	}
	return key[:32]
}

func decode(miwen string) string {
	s, _ := hex.DecodeString(miwen)
	x, _ := base64.StdEncoding.DecodeString(string(s))
	k, _ := base64.StdEncoding.DecodeString(genKey(KEY))
	str, _ := tripleDesDecrypt(x, k)
	return string(str)
}

// 3DES дешифрование
func tripleDesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, []byte(IV))
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

// 3DES шифрование
func TripleDesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	origData = PKCS5Padding(origData, block.BlockSize())
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, []byte(IV))
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// удаляем последний байт для разгрузки
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
