package logger

import (
	"testing"
	//"github.com/stretchr/testify/assert"

	"github.com/sirupsen/logrus"
)

func TestDebug(t *testing.T) {
	SetLogLevel(logrus.InfoLevel)
	req := []byte{}
	res := []byte{}
	InfoWithFields("Error", Fields{
		"req": string(req),
		"res": string(res),
	})
}
