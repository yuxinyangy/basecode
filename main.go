package main

import (
	"BaseCode/base58"
	"BaseCode/base64"
	"fmt"
)

func main() {
	str:="a"
	cipherText := base64.Base64EncodeString(str)
	fmt.Println("base64编码:",cipherText)
	fmt.Println("base64解码:",base64.Base64DecodeString(cipherText))
	cipherText1:=base58.Base58Encode([]byte(str))
	fmt.Println("base58编码:",string(cipherText1))
	fmt.Println("base58解码:",string(base58.Base58Decode(cipherText1)))
}
