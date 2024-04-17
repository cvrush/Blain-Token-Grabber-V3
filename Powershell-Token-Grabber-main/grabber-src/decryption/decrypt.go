// decrypt.go
package decryption

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"os"

	"github.com/zavla/dpapi"
)

func GetMasterKey(path string) []byte {
	data, _ := os.ReadFile(path)

	var LocalStateJson struct {
		OsCrypt struct {
			EncryptedKey string `json:"encrypted_key"`
		} `json:"os_crypt"`
	}

	_ = json.Unmarshal(data, &LocalStateJson)

	EncryptedSecretKey, _ := base64.StdEncoding.DecodeString(LocalStateJson.OsCrypt.EncryptedKey)

	secretKey := EncryptedSecretKey[5:]
	DecryptedSecretKey, _ := dpapi.Decrypt(secretKey)

	return DecryptedSecretKey
}

func DecryptPassword(buff []byte, masterKey []byte) (string, error) {
	if len(buff) < 15 {
		return "", errors.New("invalid buffer length")
	}

	iv := buff[3:15]
	payload := buff[15:]

	block, err := aes.NewCipher(masterKey)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	decryptedPass, err := gcm.Open(nil, iv, payload, nil)
	if err != nil {
		return "", err
	}

	return string(decryptedPass), nil
}
