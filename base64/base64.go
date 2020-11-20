package base64

import "encoding/base64"

func Base64EncodeString(str string)string  {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func Base64DecodeString(str string) string {
	result,_ := base64.StdEncoding.DecodeString(str)
	return string(result)
}
