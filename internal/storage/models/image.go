package models

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type ImageUnit struct {
	ID          int
	Payload     io.Reader
	PayloadName string
	PayloadSize int64
}

func (i *ImageUnit) GenerateObjectName() string {
	t := time.Now()
	formatted := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	return fmt.Sprintf(
		"%s/%s.%s",
		strconv.Itoa(i.ID),
		formatted,
		"png")
}

func (i *ImageUnit) GetObjectURL(baseUrl, bucketName, filepath string) string {
	return fmt.Sprintf(
		"%s/%s/%s",
		baseUrl,
		bucketName,
		filepath,
	)
}
