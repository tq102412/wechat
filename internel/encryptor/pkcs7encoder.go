package encryptor

import "errors"

const BlockSize = 32

// PKCS7Padding 补位
func PKCS7Padding(msg []byte) []byte {
	msgSize := len(msg)

	// 要补的位数
	amountToPad := BlockSize - (msgSize % BlockSize)

	// 0表示要补满
	if 0 == amountToPad {
		amountToPad = BlockSize
	}

	for i := 0; i < amountToPad; i++ {
		msg = append(msg, byte(amountToPad))
	}

	return msg
}

//PKCS7UnPadding 删除补位
func PKCS7UnPadding(msg []byte) ([]byte, error) {
	length := len(msg)
	lastIndex := length - 1

	if 0 > lastIndex {
		return nil, errors.New("Msg Length Error")
	}

	pad := int(msg[lastIndex])

	if pad < 1 || pad > BlockSize {
		pad = 0
	}

	newMsgLength := length - pad

	newMsg := msg[:newMsgLength]

	return newMsg, nil
}
