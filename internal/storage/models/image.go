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

func (i *ImageUnit) GenerateObjectName(bucketName string) string {
	t := time.Now()
	formatted := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	return fmt.Sprintf(
		"%s/%s/%s.%s",
		bucketName,
		strconv.Itoa(i.ID),
		formatted,
		"png")
}
