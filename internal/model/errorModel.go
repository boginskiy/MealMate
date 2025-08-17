package model

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Alerter interface {
	ToMarshal() ([]byte, error)
	ToString() (string, error)
	PreparBody(*http.Request) []byte
}

type Alert struct {
	Alert     string
	Code      int
	Message   string
	Path      string
	Timestamp time.Time
}

func NewAlert(alert string, code int, mess string, path string) *Alert {
	return &Alert{
		Alert:     alert,
		Code:      code,
		Message:   mess,
		Path:      path,
		Timestamp: time.Now(),
	}
}

func (e *Alert) ToString() (string, error) {
	slByte, err := json.Marshal(e)
	return string(slByte), err
}

func (e *Alert) ToMarshal() ([]byte, error) {
	return json.Marshal(e)
}

func (e *Alert) PreparBody(req *http.Request) []byte {
	tmpBody, err := e.ToMarshal()

	if err != nil {
		// Если не удалось сериализовать Error, пишем os.Stdout и короткое сообщение
		tmpStr := fmt.Sprintf(`{"%v message":"%v"}`, e.Alert, err)
		fmt.Fprintf(os.Stdout, "%v\n", tmpStr)
		tmpBody = []byte(tmpStr)
	}
	return tmpBody
}
