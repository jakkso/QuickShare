package files

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"path"
)

func genUUID() (string, error) {
	u := make([]byte, 16)
	_, err := rand.Read(u)
	if err != nil {
		return "", errors.New("rand read failure")
	}
	// UUID v4 compliant wizardry
	u[8] = (u[8] | 0x80) & 0xBF
	u[6] = (u[6] | 0x40) & 0x4F

	return hex.EncodeToString(u), nil
}

func GenerateFilename(dir string, ext string) (string, error) {
	uuid, err := genUUID()
	if err != nil {
		return "", err
	}
	return path.Join(dir, uuid+ext), nil
}