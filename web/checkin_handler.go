package web

import (
	"bytes"
	"encoding/base64"
	"fmt"

	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

func generateCode() (string, error) {
	code, err := qrcode.New("http://localhost:1323/event")
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	writeCloser := writeCloserWrapper{Writer: &buf}

	writer := standard.NewWithWriter(writeCloser, standard.WithBuiltinImageEncoder(standard.PNG_FORMAT))

	err = code.Save(writer)
	if err != nil {
		return "", nil
	}

	encoded := base64.StdEncoding.EncodeToString(buf.Bytes())
	return fmt.Sprintf("data:image/png;base64,%s", encoded), nil
}

