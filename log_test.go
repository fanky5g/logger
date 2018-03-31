package logger

import (
	"io/ioutil"
	"os"
	"testing"
	//"github.com/stretchr/testify/assert"
)

func TestDebug(t *testing.T) {
	file, _ := ioutil.TempFile("", "")
	logger := New("", file.Name())
	defer os.Remove(file.Name())
	req := []byte{}
	res := []byte{}
	logger.InfoMode().InfoWithFields("Paysenger Error", Fields{
		"req": string(req),
		"res": string(res),
	})
}
