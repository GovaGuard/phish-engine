package mail

import (
	"encoding/base64"
	"os"
)

type Attachement struct {
	Name          string
	Base64Content []byte
}

func NewAttachement(file string) (Attachement, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return Attachement{}, err
	}

	b := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(b, data)

	return Attachement{Name: file, Base64Content: b}, nil
}
