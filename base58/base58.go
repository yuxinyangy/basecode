package base58

import (
	"BaseCode/utils"
	"bytes"
	"math/big"
)

var base58Alphabets  = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

func Base58Encode(input []byte) []byte {
	var result []byte
	x := big.NewInt(0).SetBytes(input)
	base := big.NewInt(int64(len(base58Alphabets)))
	zero := big.NewInt(0)
	mod := &big.Int{}
	for x.Cmp(zero) != 0 {
		x.DivMod(x,base,mod)
		result = append(result,base58Alphabets[mod.Int64()])
	}
	if input[0]==0x00 {
		result = append(result,base58Alphabets[0])
	}
	utils.ReverseBytes(result)

	return result
}

func Base58Decode(input []byte) []byte  {
	result := big.NewInt(0)
	for _,b:=range input{
		charIndex := bytes.IndexByte(base58Alphabets,b)
		result.Mul(result,big.NewInt(58))
		result.Add(result,big.NewInt(int64(charIndex)))
	}
	decoded := result.Bytes()
	if input[0] == base58Alphabets[0]  {
		decoded = append([]byte{0x00},decoded...)
	}
	return decoded
}
