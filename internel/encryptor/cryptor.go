package encryptor

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"open-platform/xmlstruct"
	"strings"
	"time"
)

// DecryptMsg 要解密的消息
type DecryptMsg struct {
	Signature   string
	Timestamp   int64
	Nonce       string
	RequestData string
}

// MessageCryptor 消息加解密对象
type MessageCryptor struct {
	token          string
	encodingAESKey string
	appId          string
	aesKey         []byte
	aesIv          []byte
}

// New 生成一个消息加解密对象
func New(token string, encodingAESKey string, appId string) (MessageCryptor, error) {

	aesKey, err := base64.StdEncoding.DecodeString(encodingAESKey + "=")

	if nil != err {
		return MessageCryptor{}, errors.New("create cryptor base64 decode fail")
	}

	aesIv := aesKey[:16]

	return MessageCryptor{
		token,
		encodingAESKey,
		appId,
		aesKey,
		aesIv,
	}, nil

}

// Encrypt 加密
func (cryptor *MessageCryptor) Encrypt(text []byte) (*xmlstruct.SendEncryptMsg, error) {

	randomStr := StrRandom(16)

	var networkBytesOrder []byte = make([]byte, 4)

	binary.BigEndian.PutUint32(networkBytesOrder, uint32(len(text)))

	bytess := make([][]byte, 4)
	bytess[0] = randomStr
	bytess[1] = networkBytesOrder
	bytess[2] = text
	bytess[3] = []byte(cryptor.appId)

	textBytes := bytes.Join(bytess, []byte(""))

	//根据key 生成密文
	block, err := aes.NewCipher(cryptor.aesKey)

	if err != nil {
		return nil, err
	}

	textBytes = PKCS7Padding(textBytes)

	blockMode := cipher.NewCBCEncrypter(block, cryptor.aesIv)
	crypted := make([]byte, len(textBytes))
	blockMode.CryptBlocks(crypted, textBytes)

	cryptedstring := base64.StdEncoding.EncodeToString(crypted)
	cryptedbyte := []byte(cryptedstring)

	signatureMsg := SignatureMsg{
		string(cryptedbyte),
		cryptor.token,
		time.Now().Unix(),
		string(randomStr),
	}

	signature, err := signatureMsg.Signature()

	if nil != err {
		return nil, err
	}

	return &xmlstruct.SendEncryptMsg{
		Encrypt:      xmlstruct.Cdata{signatureMsg.EncryptMsg},
		MsgSignature: xmlstruct.Cdata{signature},
		TimeStamp:    signatureMsg.Timestamp,
		Nonce:        xmlstruct.Cdata{string(randomStr)},
	}, nil

}

//Decrypt 解密
func (cryptor *MessageCryptor) Decrypt(decryptMsg *DecryptMsg) ([]byte, error) {

	err := CheckDecryptMsg(cryptor, decryptMsg)

	if nil != err {
		return nil, err
	}

	// 解密
	ciphertextDec, err := base64.StdEncoding.DecodeString(decryptMsg.RequestData)

	if nil != err {
		return nil, err
	} else {

		decryptedMsg, err := decrypt(cryptor, ciphertextDec)

		if nil != err {
			return nil, err
		} else {
			//if "true" == envconfig.Env("APP_DEBUG") {
			//	log.Println(string(decryptedMsg))
			//}
			return decryptedMsg, err
		}

	}

}

// CheckDecryptMsg 检测解密的消息是否正确
func CheckDecryptMsg(cryptor *MessageCryptor, decryptMsg *DecryptMsg) error {

	if 43 != len(cryptor.encodingAESKey) {
		return errors.New("encodingAesKey 非法")
	}

	//验证安全签名
	signatureMsg := SignatureMsg{
		decryptMsg.RequestData,
		cryptor.token,
		decryptMsg.Timestamp,
		decryptMsg.Nonce,
	}

	signature, err := signatureMsg.Signature()

	if nil != err {
		return err
	}

	if 0 != strings.Compare(signature, decryptMsg.Signature) {
		return errors.New("签名效验失败!")
	} else {
		return nil
	}

}

// decrypt 解密内部实现
func decrypt(cryptor *MessageCryptor, decryptedMsg []byte) ([]byte, error) {

	block, err := aes.NewCipher(cryptor.aesKey)

	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, cryptor.aesIv)

	origData := make([]byte, len(decryptedMsg))

	blockMode.CryptBlocks(origData, decryptedMsg)

	origData, err = PKCS7UnPadding(origData)

	if nil != err {
		return nil, err
	}

	if len(origData) < 16 {
		return nil, errors.New("解密失败长度错误！")
	}

	convInt := binary.BigEndian.Uint32(origData[16:20])

	return origData[20 : convInt+20], nil

}
