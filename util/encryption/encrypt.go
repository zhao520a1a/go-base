package encryption

import (
	"context"
	"encoding/base64"
	"fmt"
	"sync"

	"github.com/gogo/protobuf/proto"
	"github.com/zhao520a1a/go-utils/errors"
)

const iv = "NfHK5a84jjJkwzff"

var cfg config

type config struct {
	sync.Mutex
	salts           map[int]string
	currSaltVersion int
}

func (c *config) setSalts(salts map[int]string) {
	c.Lock()
	defer c.Unlock()
	c.salts = salts

	var maxVersion int
	for v := range salts {
		if v > maxVersion {
			maxVersion = v
		}
	}
	c.currSaltVersion = maxVersion
	return
}

func (c *config) getSalt(v int) string {
	c.Lock()
	defer c.Unlock()

	return c.salts[v]
}

func (c *config) getCurrentSalt() string {
	c.Lock()
	defer c.Unlock()

	return c.salts[c.currSaltVersion]
}

func (c *config) getCurrentSaltVersion() int {
	c.Lock()
	defer c.Unlock()

	return c.currSaltVersion
}

func SetConfig(salts map[int]string) {
	cfg.setSalts(salts)
}

func EncryptText(ctx context.Context, text string) (cipherText string, err error) {
	op := errors.Op("EncryptText")

	version := cfg.getCurrentSaltVersion()
	salt := cfg.getCurrentSalt()
	if salt == "" {
		err = fmt.Errorf("salt not found")
		return
	}

	if text == "" {
		return "", nil
	}

	// TODO scrypto
	var sign []byte
	//sign, err := scrypto.CBCPKCS5PaddingAesEncrypt([]byte(salt), []byte(iv), []byte(text))
	//if err != nil {
	//	err = errors.E(op, fmt.Errorf("err %v text %s salt %s", err, text, salt))
	//	return
	//}

	info := &EncryptedInfo{
		Version: int32(version),
		Sign:    sign,
	}
	// 采用 protobuf 转成一个字节流，因其以高效的二进制方式存储，结果数据长度会较小
	body, err := proto.Marshal(info)
	if err != nil {
		err = errors.E(op, fmt.Errorf("protobuf marshal err %v", err))
		return
	}

	cipherText = base64.StdEncoding.EncodeToString(body)
	return
}

func DecryptText(ctx context.Context, cipherText string) (text string, err error) {
	op := errors.Op("DecryptText")

	cipherBytes, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		err = errors.E(op, fmt.Errorf("decode err %v", err))
		return
	}

	var info EncryptedInfo
	err = proto.Unmarshal(cipherBytes, &info)
	if err != nil {
		err = errors.E(op, fmt.Errorf("protobuf unmarshal err %v", err))
		return
	}

	salt := cfg.getSalt(int(info.Version))
	if salt == "" {
		err = errors.E(op, fmt.Errorf("salt not found for version %d", info.Version))
		return
	}

	// TODO scrypto
	var textBytes string
	//textBytes, err := scrypto.CBCPKCS5PaddingAesDecrypt([]byte(salt), []byte(iv), info.Sign)
	//if err != nil {
	//	err = errors.E(op, fmt.Errorf("aes decrypt err %v", err))
	//	return
	//}

	text = string(textBytes)
	return
}
