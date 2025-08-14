package model

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Errorer interface {
	ToMarshal() ([]byte, error)
	ToString() (string, error)
	PreparBody(*http.Request) []byte
}

type ErrorWarn struct {
	Code      int
	Message   string
	Path      string
	Timestamp time.Time
}

func NewErrorWarn(code int, mess string, path string) *ErrorWarn {
	return &ErrorWarn{
		Code:      code,
		Message:   mess,
		Path:      path,
		Timestamp: time.Now(),
	}
}

func (e *ErrorWarn) ToString() (string, error) {
	slByte, err := json.Marshal(e)
	return string(slByte), err
}

func (e *ErrorWarn) ToMarshal() ([]byte, error) {
	return json.Marshal(e)
}

func (e *ErrorWarn) PreparBody(req *http.Request, tpy string) []byte {
	tmpBody, err := e.ToMarshal()

	if err != nil {
		// Если не удалось сериализовать Error, пишем os.Stdout и короткое сообщение
		tmpStr := fmt.Sprintf(`{"%v message":"%v"}`, tpy, err)
		fmt.Fprintf(os.Stdout, "%v\n", tmpStr)
		tmpBody = []byte(tmpStr)
	}
	return tmpBody
}
