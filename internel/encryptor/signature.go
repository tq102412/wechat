package encryptor

import (
	"crypto/sha1"
	"encoding/hex"
	"sort"
	"strconv"
	"strings"
)

//SignatureMsg 要签名的数据结构
type SignatureMsg struct {
	EncryptMsg string
	Token      string
	Timestamp  int64
	Nonce      string
}

//Signature sha1对数据进行签名
func (signatureMsg SignatureMsg) Signature() (string, error) {
	// 转为字符串切片
	signatureSlice := []string{
		signatureMsg.EncryptMsg,
		signatureMsg.Token,
		strconv.FormatInt(signatureMsg.Timestamp, 10),
		signatureMsg.Nonce,
	}

	// 对字符串切片进行排序
	sort.Strings(signatureSlice)

	// 拼接字符串
	signatureMsgStr := strings.Join(signatureSlice, "")

	// 初始化sha
	sha := sha1.New()

	// 把加密字符串填充到sha中
	_, err := sha.Write([]byte(signatureMsgStr))

	// 添加字符串失败
	if nil != err {
		return "", err
	}

	// 生成hash值
	result := sha.Sum(nil)

	encodedStr := hex.EncodeToString(result)

	return encodedStr, nil
}
