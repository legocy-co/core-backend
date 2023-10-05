package helpers

import (
	"bytes"
	"errors"
	"io"
	"os"
	"path/filepath"
)

func FileExists(fp string) bool {
	_, err := os.Stat(fp)
	return !errors.Is(err, os.ErrNotExist)
}

func StreamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}

func ByteToStream(data []byte) io.Reader {
	return bytes.NewReader(data)
}

func GetConfigFilepath(fp string) string {
	cwd, _ := os.Getwd()
	return filepath.Dir(cwd) + "/" + filepath.Base(cwd) + fp
}
